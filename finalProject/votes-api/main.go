package main

import (
	"votes-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/votes", handlers.CreateVote)
		v1.GET("/votes", handlers.GetAllVotes)
		v1.GET("/votes/:id", handlers.GetVoteByID)
		v1.DELETE("/votes/:id", handlers.DeleteVote)
	}

	r.Run(":8080")
}
