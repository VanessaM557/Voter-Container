package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Poll struct {
	ID       string   `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

var polls = []Poll{}  // Mock database for our demo

// CreatePoll - creates a new poll
func CreatePoll(c *gin.Context) {
	var poll Poll
	if err := c.ShouldBindJSON(&poll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	polls = append(polls, poll)
	c.JSON(http.StatusOK, poll)
}

// GetAllPolls - retrieves all polls
func GetAllPolls(c *gin.Context) {
	c.JSON(http.StatusOK, polls)
}

// GetPollByID - retrieves a single poll by ID
func GetPollByID(c *gin.Context) {
	id := c.Param("id")

	for _, p := range polls {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
}

// UpdatePoll - updates a poll by ID
func UpdatePoll(c *gin.Context) {
	var updatedPoll Poll
	id := c.Param("id")

	for i, p := range polls {
		if p.ID == id {
			if err := c.ShouldBindJSON(&updatedPoll); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			polls[i] = updatedPoll  // Update the poll
			c.JSON(http.StatusOK, updatedPoll)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
}

// DeletePoll - deletes a poll by ID
func DeletePoll(c *gin.Context) {
	id := c.Param("id")

	for i, p := range polls {
		if p.ID == id {
			polls = append(polls[:i], polls[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Poll deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
}
