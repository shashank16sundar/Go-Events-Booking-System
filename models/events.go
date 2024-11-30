package models

import (
	"errors"
	"time"

	"shnk.com/eventx/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int64
}

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

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?,description = ?,location = ?,date_time = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	if err != nil {
		return err
	}

	return nil
}

func (event Event) Delete() error {
	query := `
	DELETE FROM events
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetEventByID(id int64) (Event, error) {
	query := `SELECT * FROM events WHERE id=?`
	res := db.DB.QueryRow(query, id)

	var queryEvent Event

	err := res.Scan(&queryEvent.ID, &queryEvent.Name, &queryEvent.Description, &queryEvent.Location, &queryEvent.DateTime, &queryEvent.UserID)

	if err != nil {
		return Event{}, err
	}
	return queryEvent, nil
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

func (e *Event) Register(userID int64) error {
	query := `INSERT INTO registrations(event_id,user_id) VALUES(?,?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userID)
	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id=? AND user_id=?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if rows == 0 {
		return errors.New("no user found registered for this event")
	}

	return err
}
