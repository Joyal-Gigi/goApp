package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goApp/models"
)

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
		//fmt.Println("Error binding JSON:", err)
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

func UpdateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve event"})
		return
	}

	var updatedEvent models.Event
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": "Request could not be parsed"})
		return
	}

	updatedEvent.ID = uint(eventId)
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update event"})
		return
	}
	c.JSON(200, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(c *gin.Context) {
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

	err = event.Delete()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete event"})
		return
	}
	c.JSON(200, gin.H{"message": "Event deleted successfully"})
}
