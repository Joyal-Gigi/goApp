package models

import (
	"time"
)

type Event struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func SaveEvent(event Event) {
	//later add database logic here
	events = append(events, event)
}

func GetAllEvents() []Event {
	return events
}
