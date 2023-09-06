package main

import (
	"log"
	"voter-api/config"
	"voter-api/handlers"
	"voter-api/data/models" 
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	
	models.InitRedis(config.Config.RedisAddr, config.Config.RedisPass) 
	
	r := gin.Default()

	r.Use(handlers.LoggingMiddleware)
        r.Use(handlers.ErrorHandlingMiddleware)
        r.Use(handlers.AuthMiddleware)
        r.Use(handlers.CORSMiddleware)

	r.GET("/voters", handlers.GetVoters)
	r.GET("/voters/:id", handlers.GetVoterByID)
	r.POST("/voters", handlers.AddVoter)
	r.PUT("/voters/:id", handlers.UpdateVoter)
	r.DELETE("/voters/:id", handlers.DeleteVoter)
	r.GET("/voters/:id/polls", handlers.GetPollsForVoter)
	r.GET("/voters/:id/polls/:pollid", handlers.GetPollForVoter)
	r.POST("/voters/:id/polls/:pollid", handlers.AddPollForVoter)
	r.PUT("/voters/:id/polls/:pollid", handlers.UpdatePollForVoter)
	r.DELETE("/voters/:id/polls/:pollid", handlers.DeletePollForVoter)
	r.GET("/voters/health", handlers.GetHealth)

	r.Run(":" + config.Config.ServerPort)
}
