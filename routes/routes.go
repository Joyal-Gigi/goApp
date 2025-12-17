package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//API endpoints
	server.GET("/ping", pingHandler)          //to check ping
	server.GET("/events", getAllEvents)       //get all events
	server.GET("/events/:id", getEvent)       //get event by id
	server.POST("/events", createEvent)       //create new event
	server.PUT("/events/:id", UpdateEvent)    //update event by id
	server.DELETE("/events/:id", deleteEvent) //delete event by id
}
