package data

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

func InitRedis(redisAddr string) {
	redisPool = &redis.Pool{
		MaxIdle:   10,
		MaxActive: 20,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddr)
		},
	}
}

func AddVoter(voter Voter) error {
	conn := redisPool.Get()
	defer conn.Close()

	voterJSON, err := json.Marshal(voter)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", fmt.Sprintf("voter:%d", voter.VoterID), voterJSON)
	return err
}

func GetVoter(voterID uint) (Voter, error) {
	conn := redisPool.Get()
	defer conn.Close()

	voterJSON, err := redis.Bytes(conn.Do("GET", fmt.Sprintf("voter:%d", voterID)))
	if err != nil {
		return Voter{}, err
	}

	var voter Voter
	err = json.Unmarshal(voterJSON, &voter)
	return voter, err
}

func UpdateVoter(voter Voter) error {
	return AddVoter(voter)
}


func DeleteVoter(voterID uint) error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", fmt.Sprintf("voter:%d", voterID))
	return err
}
