package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"your_project_path/redis"
)

type HealthCheck struct {
	Status string `json:"status"`
	Redis  string `json:"redis"`
}

func GetHealth(c *gin.Context) {
	health := HealthCheck{
		Status: "UP",
	}

	_, err := redis.Rdb.Ping(redis.Ctx).Result()
	if err != nil {
		health.Status = "DOWN"
		health.Redis = "DOWN"
	} else {
		health.Redis = "UP"
	}

	if health.Status == "UP" {
		c.JSON(http.StatusOK, health)
	} else {
		c.JSON(http.StatusInternalServerError, health)
	}
}
