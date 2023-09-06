package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Vote struct {
	ID     string `json:"id"`
	PollID string `json:"pollId"`
	VoterID string `json:"voterId"`
	Choice string `json:"choice"`
}

var votes = []Vote{}  // This is a mock database for our demo

// CreateVote - adds a new vote
func CreateVote(c *gin.Context) {
	var vote Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	votes = append(votes, vote)
	c.JSON(http.StatusOK, vote)
}

// GetAllVotes - retrieve all votes
func GetAllVotes(c *gin.Context) {
	c.JSON(http.StatusOK, votes)
}

// GetVoteByID - retrieve a single vote by ID
func GetVoteByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range votes {
		if v.ID == id {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Vote not found"})
}

// DeleteVote - deletes a vote by ID
func DeleteVote(c *gin.Context) {
	id := c.Param("id")

	for i, v := range votes {
		if v.ID == id {
			votes = append(votes[:i], votes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Vote deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Vote not found"})
}
