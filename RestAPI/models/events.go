package models

import (
	"time"

	"auth.com/auth/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:required`
	Description string    `binding:required`
	Location    string    `binding:required`
	DataTime    time.Time `binding:required`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events (name , description , location ,dataTime , user_id) 
	VALUES(?,?,?,?,? )`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DataTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	result, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var events []Event
	for result.Next() {
		var event Event
		err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DataTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE ID = ?`
	result := db.DB.QueryRow(query, id)
	var event Event
	err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DataTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) DeleteEvent() error {
	query := `DELETE  FROM events WHERE ID = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)

	return err
}

func (e Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET name = ? , description = ? , location = ? , dataTime = ?
	WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DataTime, e.ID)

	return err

}

func (e Event) Register(userId int64) error {
	query := `INSERT INTO registrations (event_id , user_id) VALUES (? , ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
