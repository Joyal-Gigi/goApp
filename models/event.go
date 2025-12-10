package models

import (
	"time"

	"github.com/goApp/db"
)

// Event represents an event entity
type Event struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	UserID      int       `json:"user_id"`
}

// SaveEvent saves a new event to the database
func SaveEvent(event Event) error {
	query := `INSERT INTO events (title, description, location, datetime, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(event.Title, event.Description, event.Location, event.Date, event.UserID)
	if err != nil {
		//fmt.Println("Error executing query:", err)
		return err
	}
	id, err := result.LastInsertId()
	event.ID = uint(id)
	return err
}

// GetAllEvents retrieves all events from the database
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

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT id, title, description, location, datetime, user_id FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.Date, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
