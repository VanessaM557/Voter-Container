package main

import (
    "github.com/gin-gonic/gin"
    "github.com/yourusername/votes-api/api/handlers"
)

func registerRoutes() *gin.Engine {
    r := gin.Default()

    // Votes
    r.POST("/votes", handlers.CreateVote)
    r.GET("/votes", handlers.GetVotes)
    r.GET("/votes/:id", handlers.GetVoteByID)
    r.DELETE("/votes/:id", handlers.DeleteVoteByID)

    return r
}
