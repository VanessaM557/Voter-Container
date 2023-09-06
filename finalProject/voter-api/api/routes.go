package api

import (
	"github.com/gin-gonic/gin"
	"voter-api/api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/voters/health", handlers.GetHealth)
  
	r.GET("/voters", handlers.GetAllVoters)
	r.GET("/voters/:id", handlers.GetVoterByID)
	r.POST("/voters/:id", handlers.AddVoter)
	r.PUT("/voters/:id", handlers.UpdateVoter)
	r.DELETE("/voters/:id", handlers.DeleteVoter)

	r.GET("/voters/:id/polls", handlers.GetPollsByVoterID)
	r.GET("/voters/:id/polls/:pollid", handlers.GetPollByVoterAndPollID)
	r.POST("/voters/:id/polls/:pollid", handlers.AddPollToVoter)
	r.PUT("/voters/:id/polls/:pollid", handlers.UpdatePollForVoter)
	r.DELETE("/voters/:id/polls/:pollid", handlers.DeletePollFromVoter)

	return r
}
