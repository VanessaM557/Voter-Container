package main
import (
	"encoding/json"
	"log"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var ctx = context.Background()

//Structs for voters
type voterPoll struct {
	PollID   uint      `json:"poll_id"`
	VoteDate time.Time `json:"vote_date"`
}

type Voter struct {
	VoterID     uint        `json:"voter_id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	VoteHistory []voterPoll `json:"vote_history"`
}

//incorportate redis
type VoterList struct {
	rdb *redis.Client
}

func NewVoterList() *VoterList {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return &VoterList{rdb: rdb}
}

func (vl *VoterList) GetAllVoters() ([]Voter, error) {
	keys, err := vl.rdb.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	
	var voters []Voter
	for _, key := range keys {
		voterData, err := vl.rdb.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		var v Voter
		err = json.Unmarshal([]byte(voterData), &v)
		if err != nil {
			return nil, err
		}
		voters = append(voters, v)
	}
	return voters, nil
}

func (vl *VoterList) GetVoterByID(id uint) (*Voter, error) {
	data, err := vl.rdb.Get(ctx, strconv.Itoa(int(id))).Result()
	if err != nil {
		return nil, err
	}
	var v Voter
	err = json.Unmarshal([]byte(data), &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (vl *VoterList) AddOrUpdateVoter(v Voter) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return vl.rdb.Set(ctx, strconv.Itoa(int(v.VoterID)), data, 0).Err()
}

func (vl *VoterList) DeleteVoterByID(id uint) error {
	return vl.rdb.Del(ctx, strconv.Itoa(int(id))).Err()
}

func main() {
	r := gin.Default()
	voterList := NewVoterList()
	defer voterList.rdb.Close()

	r.GET("/voters", func(c *gin.Context) {
		voters, err := voterList.GetAllVoters()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, voters)
	})

	r.GET("/voters/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid Voter ID"})
			return
		}
		voter, err := voterList.GetVoterByID(uint(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Voter Not Found"})
			return
		}
		c.JSON(200, voter)
	})

	r.POST("/voters/:id", func(c *gin.Context) {
		var v Voter
		if err := c.BindJSON(&v); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := voterList.AddOrUpdateVoter(v); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, v)
	})

	r.PUT("/voters/:id", func(c *gin.Context) {
		var v Voter
		if err := c.BindJSON(&v); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := voterList.AddOrUpdateVoter(v); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, v)
	})

	r.DELETE("/voters/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid Voter ID"})
			return
		}
		if err := voterList.DeleteVoterByID(uint(id)); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "Voter Deleted"})
	})

	r.GET("/voters/:id/polls", func(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Voter ID"})
		return
	}
	voter, err := voterList.GetVoterByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Voter Not Found"})
		return
	}
	c.JSON(200, voter.VoteHistory)
})

r.GET("/voters/:id/polls/:pollid", func(c *gin.Context) {
	voterID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Voter ID"})
		return
	}
	pollID, err := strconv.Atoi(c.Param("pollid"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Poll ID"})
		return
	}
	voter, err := voterList.GetVoterByID(uint(voterID))
	if err != nil {
		c.JSON(404, gin.H{"error": "Voter Not Found"})
		return
	}
	var specificPoll *voterPoll
	for _, poll := range voter.VoteHistory {
		if poll.PollID == uint(pollID) {
			specificPoll = &poll
			break
		}
	}
	if specificPoll == nil {
		c.JSON(404, gin.H{"error": "Poll Not Found for Given Voter"})
		return
	}
	c.JSON(200, specificPoll)
})

r.POST("/voters/:id/polls/:pollid", func(c *gin.Context) {
	var poll voterPoll
	if err := c.BindJSON(&poll); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	voterID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Voter ID"})
		return
	}
	voter, err := voterList.GetVoterByID(uint(voterID))
	if err != nil {
		c.JSON(404, gin.H{"error": "Voter Not Found"})
		return
	}
	voter.VoteHistory = append(voter.VoteHistory, poll)
	if err := voterList.AddOrUpdateVoter(*voter); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, poll)
})
	
r.PUT("/voters/:id/polls/:pollid", func(c *gin.Context) {
	voterID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Voter ID"})
		return
	}
	pollID, err := strconv.Atoi(c.Param("pollid"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Poll ID"})
		return
	}
	voter, err := voterList.GetVoterByID(uint(voterID))
	if err != nil {
		c.JSON(404, gin.H{"error": "Voter Not Found"})
		return
	}

	var updatedPoll voterPoll
	if err := c.BindJSON(&updatedPoll); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	pollFound := false
	for i, poll := range voter.VoteHistory {
		if poll.PollID == uint(pollID) {
			voter.VoteHistory[i] = updatedPoll
			pollFound = true
			break
		}
	}
	if !pollFound {
		c.JSON(404, gin.H{"error": "Poll Not Found for Given Voter"})
		return
	}

	if err := voterList.AddOrUpdateVoter(*voter); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "Poll Updated Successfully"})
})

r.DELETE("/voters/:id/polls/:pollid", func(c *gin.Context) {
	voterID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Voter ID"})
		return
	}
	pollID, err := strconv.Atoi(c.Param("pollid"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Poll ID"})
		return
	}
	voter, err := voterList.GetVoterByID(uint(voterID))
	if err != nil {
		c.JSON(404, gin.H{"error": "Voter Not Found"})
		return
	}
	pollIndex := -1
	for i, poll := range voter.VoteHistory {
		if poll.PollID == uint(pollID) {
			pollIndex = i
			break
		}
	}
	if pollIndex == -1 {
		c.JSON(404, gin.H{"error": "Poll Not Found for Given Voter"})
		return
	}
	voter.VoteHistory = append(voter.VoteHistory[:pollIndex], voter.VoteHistory[pollIndex+1:]...)
	if err := voterList.AddOrUpdateVoter(*voter); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "Poll Deleted Successfully"})
})
	r.Run(":8080")
}
