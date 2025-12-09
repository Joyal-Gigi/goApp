package models

import (
	"time"

	"github.com/goApp/db"
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

func SaveEvent(event Event) error {
	query := `INSERT INTO events (title, description, location, date, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(event.Title, event.Description, event.Location, event.Date, event.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	event.ID = uint(id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `Select * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.Date, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
