package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()

	// endpoint 1 - hello world
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// end point 2 - display client IP
	router.GET("/ip", func(c *gin.Context) {
		//fmt.Printf("ClientIP: %s\n", c.ClientIP())
		c.JSON(200, gin.H{
			"Client IP Address": c.ClientIP(),
		})
	})

	router.Use(gin.LoggerWithFormatter(customLogger))

	router.Run("localhost:8080")
}

func customLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
