package handlers

import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	log.Printf("%s - %s - %d - %s",
		c.Request.Method,
		c.Request.URL.Path,
		c.Writer.Status(),
		time.Since(startTime),
	)
}

func ErrorHandlingMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(c.Writer.Status(), gin.H{"errors": c.Errors.Errors()})
	}
}

func AuthenticationMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if token != "expected_token" {
		c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
		return
	}

	c.Next()
}
