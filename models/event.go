package models

import (
	"time"

	"shnk.com/eventx/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

var Events = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO 
	events(name,description,location,date_time,user_id)
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var events []Event
	for res.Next() {
		var event Event
		err := res.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
