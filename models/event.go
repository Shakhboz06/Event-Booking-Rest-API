package models

import (
	"event-booking-rest-api/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	Datetime    time.Time `binding: "required"`
	UserID      int64

}

var events = []Event{}

func (event *Event) Save() error {
	query := `INSERT INTO Events(name, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stm, err := db.DB.Prepare(query)

	if err != nil{
		panic("Cannot be inserted: " + err.Error())
	}

	defer stm.Close()

	res, err := stm.Exec(event.Name, event.Description, event.Location, event.Datetime, event.UserID)

	if err != nil{
		return err
	}
	
	id, err := res.LastInsertId()
	event.Id = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM Events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventbyID(id int64) (*Event, error){
	query := "SELECT * FROM Events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)

	if err != nil{
		return nil, err
	}
	
	return &event, nil
}

func(e Event) Update() error {
	query := `
	UPDATE Events
	SET name = ?, description = ?, location = ?, datetime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.Id)

	return err
}

func(e Event) Delete() error{
	query := `DELETE FROM Events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id)

	return err
}


func ( e Event) Register(userId int64) error{
	query := `INSERT INTO Registrations (user_id, event_id) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)

	return err
}

func ( e Event) Cancel(userId int64) error{
	query := `DELETE FROM Registrations WHERE user_id = ? AND event_id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)

	return err
}
