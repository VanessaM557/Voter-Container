package main

import (
	"polls-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/polls", handlers.CreatePoll)
		v1.GET("/polls", handlers.GetAllPolls)
		v1.GET("/polls/:id", handlers.GetPollByID)
		v1.PUT("/polls/:id", handlers.UpdatePoll)
		v1.DELETE("/polls/:id", handlers.DeletePoll)
	}

	r.Run(":8080")
}
