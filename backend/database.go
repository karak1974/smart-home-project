package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("USER"),
		os.Getenv("PASS"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		"smarthome")
)

func getDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}

// AddRecord adds a log record about state of a lamp
// We don't give the ID due the database will create it
func AddRecord(eventLog EventLog) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO event_logs (lamp, date, status) VALUES (? ,?, ?)",
		eventLog.Lamp,
		eventLog.Date,
		eventLog.Status)
	return err
}

// GetRecordById return a record where the provided ID appears
func GetRecordById(recordId int) (EventLog, error) {
	db, err := getDB()
	if err != nil {
		return EventLog{}, err
	}

	events, err := db.Query("SELECT * FROM event_logs WHERE id=? LIMIT 1", recordId)
	if err != nil {
		return EventLog{}, err
	}

	var res EventLog
	var id int
	var lamp, date string
	var status bool

	if err = events.Scan(&id, &lamp, &date, &status); err != nil {
		return EventLog{}, err
	}

	res.Id = id
	res.Lamp = lamp
	res.Date = date
	res.Status = status

	return res, nil
}

// GetRecordByLamp return a record where the provided Lamp name
func GetRecordByLamp(recordLamp string) ([]EventLog, error) {
	db, err := getDB()
	if err != nil {
		return []EventLog{}, err
	}

	events, err := db.Query("SELECT * FROM event_logs WHERE lamp=?", recordLamp)
	if err != nil {
		return []EventLog{}, err
	}

	var tmp EventLog
	var res []EventLog

	for events.Next() {
		var id int
		var lamp, date string
		var status bool
		err = events.Scan(&id, &lamp, &date, &status)
		if err != nil {
			return []EventLog{}, err
		}
		tmp.Id = id
		tmp.Lamp = lamp
		tmp.Date = date
		tmp.Status = status
		res = append(res, tmp)
	}

	return res, nil
}

// GetEventByDate ([]EventLog, error)
func GetRecordByDate(startDate, endDate string) ([]EventLog, error) {
	db, err := getDB()
	if err != nil {
		return []EventLog{}, err
	}

	events, err := db.Query("SELECT * FROM event_logs WHERE date BETWEEN ? AND ?", startDate, endDate)
	if err != nil {
		return []EventLog{}, err
	}

	var tmp EventLog
	var res []EventLog

	for events.Next() {
		var id int
		var lamp, date string
		var status bool

		if err := events.Scan(&id, &lamp, &date, &status); err != nil {
			return []EventLog{}, err
		}
		tmp.Id = id
		tmp.Lamp = lamp
		tmp.Date = date
		tmp.Status = status
		res = append(res, tmp)
	}

	return res, nil
}
