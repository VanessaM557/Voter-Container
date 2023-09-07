package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"your_project_path/redis"
	"your_project_path/models"
	"strconv"
)

func CreatePoll(c *gin.Context) {
	var poll models.Poll
	
	if err := c.ShouldBindJSON(&poll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pollID := strconv.Itoa(rand.Int()) 
	poll.ID = pollID
	pollJSON, _ := json.Marshal(poll)
	
	err := redis.Rdb.Set(redis.Ctx, pollID, pollJSON, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the poll"})
		return
	}

	c.JSON(http.StatusCreated, poll)
}

func GetPoll(c *gin.Context) {
	id := c.Param("id")

	val, err := redis.Rdb.Get(redis.Ctx, id).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
		return
	}

	var poll models.Poll
	err = json.Unmarshal([]byte(val), &poll)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse poll"})
		return
	}

	selfLink := models.Link{Rel: "self", Href: "/polls/" + id}
	voteLink := models.Link{Rel: "vote", Href: "/polls/" + id + "/vote"}
	poll.Links = []models.Link{selfLink, voteLink}
	c.JSON(http.StatusOK, poll)
}

func UpdatePoll(c *gin.Context) {
	id := c.Param("id")

	var updatedPoll models.Poll
	if err := c.ShouldBindJSON(&updatedPoll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPollJSON, _ := json.Marshal(updatedPoll)
	
	err := redis.Rdb.Set(redis.Ctx, id, updatedPollJSON, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update the poll"})
		return
	}

	c.JSON(http.StatusOK, updatedPoll)
}

func DeletePoll(c *gin.Context) {
	id := c.Param("id")

	err := redis.Rdb.Del(redis.Ctx, id).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete poll"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Poll deleted successfully"})
}

