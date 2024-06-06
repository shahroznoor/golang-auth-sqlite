package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to Database")
		// log.Fatalf("could not connect to Database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	if err = DB.Ping(); err != nil {
		log.Fatalf("could not ping the Database: %v", err)
	}

	createTables()
}

// createTables ensures the required tables are created
func createTables() {
	createEventTables := `
	CREATE TABLE IF NOT EXISTS events (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL,
    description TEXT NOT NULL,
    location    TEXT NOT NULL,
    dataTime    DATETIME NOT NULL,
    user_id     INTEGER,
    FOREIGN KEY(user_id) REFERENCES user(id)
);`
	_, err := DB.Exec(createEventTables)
	if err != nil {
		panic("create event table query error")
		// log.Fatalf("create event table query error: %v", err)
	}

	createUserTables := `
	CREATE TABLE IF NOT EXISTS user (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		name        TEXT NOT NULL,
		email		TEXT NOT NULL UNIQUE,
		password	TEXT NOT NULL
	)`

	_, err = DB.Exec(createUserTables)
	if err != nil {
		panic("create user table query error")
		// log.Fatalf("create event table query error: %v", err)
	}

	createRegistrationTables := `
	CREATE TABLE IF NOT EXISTS registrations (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id    INTEGER,
		user_id     INTEGER,
		FOREIGN KEY(event_id) REFERENCES event(id),
		FOREIGN KEY(user_id) REFERENCES user(id)
	)`

	_, err = DB.Exec(createRegistrationTables)
	if err != nil {
		panic("create registrer table query error")
		// log.Fatalf("create event table query error: %v", err)
	}

	fmt.Println("Tables created successfully!")
}
