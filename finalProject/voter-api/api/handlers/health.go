package handlers

import (
    "encoding/json"
    "net/http"
    "sync"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/gomodule/redigo/redis"
)

var (
	bootTime            = time.Now()
	totalAPICalls       int64
	totalAPICallsErrors int64
	mu                  sync.Mutex 
	redisPool *redis.Pool
)

type Health struct {
	Uptime             string `json:"uptime"`
	TotalAPICalls      int64  `json:"total_api_calls"`
	TotalAPICallsError int64  `json:"total_api_calls_errors"`
}

func InitRedis(redisAddr string) {
	redisPool = &redis.Pool{
		MaxIdle:   10,
		MaxActive: 20, // adjust to your needs
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddr)
		},
	}
}

func checkRedisHealth() error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")
	return err
}

func incrementAPICalls() {
	mu.Lock()
	defer mu.Unlock()
	totalAPICalls++
}

func incrementAPICallsErrors() {
	mu.Lock()
	defer mu.Unlock()
	totalAPICallsErrors++
}

func GetHealth(c *gin.Context) {
	incrementAPICalls()
	err := checkRedisHealth()
	if err != nil {
		incrementAPICallsErrors()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis is down"})
		return
	}

	uptime := time.Since(bootTime).String()

	health := Health{
		Uptime:             uptime,
		TotalAPICalls:      totalAPICalls,
		TotalAPICallsError: totalAPICallsErrors,
	}

	c.JSON(http.StatusOK, health)
}

func AddVoter(c *gin.Context) {
    incrementAPICalls()
    
    var newVoter Voter
    if err := c.ShouldBindJSON(&newVoter); err != nil {
        incrementAPICallsErrors()
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Voter added successfully!"})
}
func GetVoterByID(c *gin.Context) {
    incrementAPICalls()

    voterID := "voter:" + c.Param("id")
    
    conn := redisPool.Get()
    defer conn.Close()
  
    exists, err := redis.Bool(conn.Do("EXISTS", voterID))
    if err != nil {
        incrementAPICallsErrors()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    if !exists {
        incrementAPICallsErrors()
        c.JSON(http.StatusNotFound, gin.H{"error": "Voter not found"})
        return
    }

    voterData, err := redis.String(conn.Do("GET", voterID))
    if err != nil {
        incrementAPICallsErrors()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch voter details"})
        return
    }

    var voter Voter
    err = json.Unmarshal([]byte(voterData), &voter)
    if err != nil {
        incrementAPICallsErrors()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process voter details"})
        return
    }

    c.JSON(http.StatusOK, voter)
}





