package models

import (
	"errors"
	"fmt"
	"time"
)

type VoterPoll struct {
	PollID   uint
	VoteDate time.Time
}

type Voter struct {
	VoterID     uint
	FirstName   string
	LastName    string
	VoteHistory []VoterPoll
}

func (v *Voter) FullName() string {
	return fmt.Sprintf("%s %s", v.FirstName, v.LastName)
}

func (v *Voter) TotalVotes() int {
	return len(v.VoteHistory)
}

func (v *Voter) HasVotedIn(pollID uint) bool {
	for _, vp := range v.VoteHistory {
		if vp.PollID == pollID {
			return true
		}
	}
	return false
}

func (v *Voter) LastVotedPoll() (*VoterPoll, error) {
	if len(v.VoteHistory) == 0 {
		return nil, errors.New("voter has no voting history")
	}
	latestVote := v.VoteHistory[0]
	for _, vp := range v.VoteHistory {
		if vp.VoteDate.After(latestVote.VoteDate) {
			latestVote = vp
		}
	}
	return &latestVote, nil
}

func (v *Voter) AddVoterPoll(vp VoterPoll) {
	v.VoteHistory = append(v.VoteHistory, vp)
}

func (v *Voter) AddMultipleVoterPolls(vps []VoterPoll) {
	v.VoteHistory = append(v.VoteHistory, vps...)
}

func (v *Voter) ClearVoteHistory() {
	v.VoteHistory = []VoterPoll{}
}

func (v *Voter) ResetVoterID(newID uint) {
	v.VoterID = newID
}
