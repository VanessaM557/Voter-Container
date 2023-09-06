package models

import (
	"errors"
	"project-root/votes-api/data" 
)

type Vote struct {
	ID       string `json:"id"`
	VoterID  string `json:"voter_id"`
	PollID   string `json:"poll_id"`
	Choice   string `json:"choice"` 
	Timestamp string `json:"timestamp"`
}

func AllVotes() ([]Vote, error) {
	return data.GetAllVotesFromRedis() 
}

func GetVoteByID(voteID string) (*Vote, error) {
	return data.GetVoteFromRedisByID(voteID)
}

func SaveVote(vote *Vote) error {
	if vote.ID == "" {
		return errors.New("vote ID is missing")
	}
	return data.SaveVoteToRedis(vote)
}

func UpdateVoteByID(voteID string, updatedVote *Vote) error {
	if updatedVote.ID == "" {
		return errors.New("updated vote ID is missing")
	}
	return data.UpdateVoteInRedisByID(voteID, updatedVote)
}

func DeleteVoteByID(voteID string) error {
	return data.DeleteVoteFromRedisByID(voteID)
}
