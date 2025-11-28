package models

import "time"

type Event struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        time.Time `json:"date"`
	UserID      int   `json:"user_id"`
}

var Events = []Event{}

func SaveEvent(event Event) {
	Events = append(Events, event)
}