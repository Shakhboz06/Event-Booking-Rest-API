package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to Database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createEventsTable()
}

func createEventsTable() {

	userTable := `
		CREATE TABLE IF NOT EXISTS Users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_name TEXT NOT NULL,
			user_surname TEXT NOT NULL,
			user_email TEXT NOT NULL UNIQUE,
			user_password TEXT NOT NULL
		)
	`
	_, err := DB.Exec(userTable)

	if err != nil {
		panic("Database user creation error: " + err.Error())
	}

	eventTable := `
		CREATE TABLE IF NOT EXISTS Events(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			datetime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES Users(id) 
		)
	`

	_, err = DB.Exec(eventTable)

	if err != nil {
		panic("Database creation error: " + err.Error())
	}

	registrationsTable := `
		CREATE TABLE IF NOT EXISTS Registrations(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(event_id) REFERENCES Events(id),
			FOREIGN KEY(user_id) REFERENCES Users(id)
		)`

	_, err = DB.Exec(registrationsTable)	

	if err != nil {
		panic("Database creation error: " + err.Error())
	}
}
