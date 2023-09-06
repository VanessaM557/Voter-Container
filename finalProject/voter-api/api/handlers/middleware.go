package handlers

import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		log.Printf("[Logger] %v %v %v %v",
			c.Request.Method,
			c.Request.URL,
			c.Writer.Status(),
			endTime.Sub(startTime))
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
    
		if len(c.Errors) > 0 {
			c.JSON(c.Writer.Status(), gin.H{
				"errors": c.Errors,
			})
		}
	}
}

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "some-secret-token" { // In real world, compare with a real token or use JWTs
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No content
			return
		}
		c.Next()
	}
}
