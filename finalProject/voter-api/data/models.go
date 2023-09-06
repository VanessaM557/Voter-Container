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

// Returns the full name of the voter.
func (v *Voter) FullName() string {
	return fmt.Sprintf("%s %s", v.FirstName, v.LastName)
}

// Returns the number of polls the voter has participated in.
func (v *Voter) TotalVotes() int {
	return len(v.VoteHistory)
}

// Check if the voter has voted in a specific poll.
func (v *Voter) HasVotedIn(pollID uint) bool {
	for _, vp := range v.VoteHistory {
		if vp.PollID == pollID {
			return true
		}
	}
	return false
}

// Get the last poll the voter has voted in, if available.
func (v *Voter) LastVotedPoll() (*VoterPoll, error) {
	if len(v.VoteHistory) == 0 {
		return nil, errors.New("voter has no voting history")
	}
	// Assuming the VoteHistory is not necessarily sorted, we find the most recent vote date.
	latestVote := v.VoteHistory[0]
	for _, vp := range v.VoteHistory {
		if vp.VoteDate.After(latestVote.VoteDate) {
			latestVote = vp
		}
	}
	return &latestVote, nil
}

// Add a new poll to the voter's voting history.
func (v *Voter) AddVoterPoll(vp VoterPoll) {
	v.VoteHistory = append(v.VoteHistory, vp)
}

// Add multiple polls at once to the voter's voting history.
func (v *Voter) AddMultipleVoterPolls(vps []VoterPoll) {
	v.VoteHistory = append(v.VoteHistory, vps...)
}

// Clear the voting history of a voter.
func (v *Voter) ClearVoteHistory() {
	v.VoteHistory = []VoterPoll{}
}

// Reset the voter's ID.
func (v *Voter) ResetVoterID(newID uint) {
	v.VoterID = newID
}
