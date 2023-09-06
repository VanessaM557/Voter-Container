package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Voter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var voters = []Voter{}  // Mock database for our demo

// RegisterVoter - registers a new voter
func RegisterVoter(c *gin.Context) {
	var voter Voter
	if err := c.ShouldBindJSON(&voter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	voters = append(voters, voter)
	c.JSON(http.StatusOK, voter)
}

// GetAllVoters - retrieves all voters
func GetAllVoters(c *gin.Context) {
	c.JSON(http.StatusOK, voters)
}

// GetVoterByID - retrieves a single voter by ID
func GetVoterByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range voters {
		if v.ID == id {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Voter not found"})
}

// DeleteVoter - deletes a voter by ID
func DeleteVoter(c *gin.Context) {
	id := c.Param("id")

	for i, v := range voters {
		if v.ID == id {
			voters = append(voters[:i], voters[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Voter deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Voter not found"})
}
