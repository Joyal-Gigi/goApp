package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goApp/db"
	"github.com/goApp/models"
)

func main() {
	server := gin.Default() //Create Gin server instance
	db.InitDB()             //Initialize Database

	//API endpoints
	server.GET("/ping", pingHandler)    //to check ping
	server.GET("/events", getAllEvents) //get all events
	server.GET("/events/:id", getEvent) //get event by id
	server.POST("/events", createEvent) //create new event

	//start server
	server.Run(":8080") //localhost:8080
	fmt.Println("Server is running on http://localhost:8080")
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getAllEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve events"})
		return
	}
	c.JSON(200, events)
}

func getEvent(c *gin.Context) {
	// Implementation for getting a single event by ID
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve event"})
		return
	}
	c.JSON(200, event)
}

func createEvent(c *gin.Context) {
	event := models.Event{}
	if err := c.ShouldBindJSON(&event); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": "Request could not be parsed"})
		return
	}
	err := models.SaveEvent(event)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save event"})
		return
	}
	c.JSON(201, gin.H{"message": "Event created", "event": event})
}
