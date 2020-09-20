package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var GitCommit string

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"commit":  GitCommit,
		})
	})
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
