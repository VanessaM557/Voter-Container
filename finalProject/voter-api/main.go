package main

import (
	"voter-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/voters", handlers.RegisterVoter)
		v1.GET("/voters", handlers.GetAllVoters)
		v1.GET("/voters/:id", handlers.GetVoterByID)
		v1.DELETE("/voters/:id", handlers.DeleteVoter)
	}

	r.Run(":8080")
}
