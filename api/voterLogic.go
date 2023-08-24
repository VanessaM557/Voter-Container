package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)



func DeleteVoterByID(appContext *AppContext, id string) error {
	err := appContext.Rdb.Del(appContext.Ctx, "voter:"+id).Err()
	if err != nil {
		return fmt.Errorf("error deleting voter: %v", err)
	}
	return nil
}

func UpdateByID(appContext *AppContext, id string, firstName, lastName string) error {
	voterJSON, err := appContext.Rdb.Get(appContext.Ctx, "voter:"+id).Result()
	if err == redis.Nil {
		return fmt.Errorf("voter cannot be found")
	} else if err != nil {
		return err
	}

	var voter Voter
	err = json.Unmarshal([]byte(voterJSON), &voter)
	if err != nil {
		return err
	}

	voter.FirstName = firstName
	voter.LastName = lastName   
	voterJSONBytes, err := json.Marshal(voter)
	if err != nil {
		return err
	}

	err = appContext.Rdb.Set(appContext.Ctx, "voter:"+id, string(voterJSONBytes), 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func CreateNewVoter(appContext *AppContext, id, firstName, lastName string) (Voter, error) {
	voter := Voter{
		ID:        id,
		FirstName: firstName, 
		LastName:  lastName,  
	}
	voterJSON, err := json.Marshal(voter)
	if err != nil {
		return Voter{}, err
	}
	err = appContext.Rdb.Set(appContext.Ctx, "voter:"+id, voterJSON, 0).Err()
	if err != nil {
		return Voter{}, err
	}
	return voter, nil
}


func GetAllVoters(appContext *AppContext) ([]Voter, error) {
	voterKeys, err := appContext.Rdb.Keys(appContext.Ctx, "voter:*").Result()
	if err != nil {
		return nil, err
	}
	voters := []Voter{}
	for _, key := range voterKeys {
		voterJSON, err := appContext.Rdb.Get(appContext.Ctx, key).Result()
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

func GetVoterByID(appContext *AppContext, id string) (Voter, error) {
	voterJSON, err := appContext.Rdb.Get(appContext.Ctx, "voter:"+id).Result()
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
