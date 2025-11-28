package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", pingHandler)

	server.Run(":8080") //localhost:8080
	fmt.Println("Server is running on http://localhost:8080")
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}