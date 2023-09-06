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

type HypermediaPoll struct {
	Poll
	Links map[string]string `json:"_links"`
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
	
	// Including hypermedia links
	hp := HypermediaPoll{
		Poll:  poll,
		Links: map[string]string{
			"self": "/v1/polls/" + poll.ID,
		},
	}
	c.JSON(http.StatusOK, hp)
}

// GetAllPolls - retrieves all polls
func GetAllPolls(c *gin.Context) {
	hypermediaPolls := []HypermediaPoll{}
	for _, p := range polls {
		hp := HypermediaPoll{
			Poll:  p,
			Links: map[string]string{
				"self": "/v1/polls/" + p.ID,
				"votes": "/v1/votes?pollId=" + p.ID,
			},
		}
		hypermediaPolls = append(hypermediaPolls, hp)
	}
	c.JSON(http.StatusOK, hypermediaPolls)
}

// GetPollByID - retrieves a single poll by ID
func GetPollByID(c *gin.Context) {
	id := c.Param("id")

	for _, p := range polls {
		if p.ID == id {
			hp := HypermediaPoll{
				Poll:  p,
				Links: map[string]string{
					"self": "/v1/polls/" + p.ID,
					"votes": "/v1/votes?pollId=" + p.ID,
					"delete": "/v1/polls/" + p.ID,
					"update": "/v1/polls/" + p.ID,
				},
			}
			c.JSON(http.StatusOK, hp)
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

			hp := HypermediaPoll{
				Poll:  updatedPoll,
				Links: map[string]string{
					"self": "/v1/polls/" + updatedPoll.ID,
					"votes": "/v1/votes?pollId=" + updatedPoll.ID,
				},
			}
			c.JSON(http.StatusOK, hp)
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
