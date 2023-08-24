package main
import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

//adding in UPDATE

func (id string, firstName, lastName string) error {
	voterIndex, exists := rdb.Get(context, "voter:"+id).Result()
	if err == redis.Nil{
		return fmt.Errorf("voter cannot be found")
	}
        
	else if err != nil {
		return err
	}

	var voter Voter

	err = json.Unmarshal([]byte(voterJSON), &voter)
	if err != nil {
		return err
	}
	
	voter.Name = firstName + " " + lastName
	voterJSON, err = json.Marshal(voter)
	
	if err != nil {
		return err
	}

	err = rdb.Set(ctx, "voter:"+id, voterJSON, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Adding in DELETE:
func DeleteVoterByID(id string) error{
	err := rdb.Del(context , "voter:"+id).Err()
	if err != nil {
		return fmt.Errorf("error deleting voter: %v", err)
	}
	return nil
}

// Function to create a new unique voter 

func CreateNewVoter(id, name string) (Voter, error){
	voter := Voter{
		ID:   id,
		Name: name,
	}
	voterJSON, err := json.Marshal(voter)
	if err != nil {
		return Voter{}, err
	}
	err = rdb.Set(context, "voter:"+id, voterJSON, 0).Err()
	if err != nil {
		return Voter{}, err
	}
	return voter, nil
}

// retrieving voters

func GetAllVoters()([]Voter, error){
	voterKeys, err := rdb.Keys(context, "voter:*").Result()
	if err != nil {
		return nil, err
	}
	voters := []Voter{}
	for _, key := range voterKeys {
		voterJSON, err := rdb.Get(context, key).Result()
		if err != nil {
			return nil, err
		}
		var voter Voter
		err = json.Unmarshal([]byte(voterJSON), &voter)
		if err != nil {
			return nil, err
		}
		voters = append(voters, voter)
	}
	return voters, nil
}

//retrieving voter by ID
func GetVoterByID(id string) (Voter, error) {
	voterJSON, err := rdb.Get(context, "voter:"+id).Result()
	if err == redis.Nil {
		return Voter{}, fmt.Errorf("voter not found")
	} else if err != nil {
		return Voter{}, err
	}
	var voter Voter
	err = json.Unmarshal([]byte(voterJSON), &voter)
	if err != nil {
		return Voter{}, err
	}
	return voter, nil
}
