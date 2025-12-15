package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goApp/db"
	"github.com/goApp/routes"
)

func main() {
	server := gin.Default() //Create Gin server instance
	db.InitDB()             //Initialize Database

	routes.RegisterRoutes(server) //Register API routes

	//start server
	server.Run(":8080") //localhost:8080
	fmt.Println("Server is running on http://localhost:8080")
}
