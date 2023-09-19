package db

import (
	"database/sql"
	"log/slog"
	"smarthome/types"
	"smarthome/vars"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() (*sql.DB, error) {
	return sql.Open("mysql", vars.ConnectionString)
}

// AddRecord adds a log record about status2 of a lamp
// We don't give the ID due the database will create it
func AddRecord(eventLog types.EventLog) (types.EventLog, error) {
	db, err := getDB()
	if err != nil {
		return types.EventLog{}, err
	}

	slog.Info("DEBUG",
		slog.String("lamp", eventLog.Lamp),
		slog.Any("date", eventLog.Date),
		slog.Bool("status", eventLog.Status))
	// Save the record in the database
	_, err = db.Exec("INSERT INTO event_logs (lamp, date, status) VALUES (?, NOW(), ?)",
		eventLog.Lamp,
		eventLog.Status)
	if err != nil {
		return types.EventLog{}, err
	}

	var res types.EventLog
	var id int
	var lamp, date string
	var status bool

	// Get the record from the database
	// This part is necessary due the database generate some values
	events := db.QueryRow("SELECT * FROM event_logs WHERE lamp=? ORDER BY id DESC", eventLog.Lamp)
	if err = events.Scan(&id, &lamp, &date, &status); err != nil {
		return types.EventLog{}, err
	}

	res.Id = id
	res.Lamp = lamp
	res.Date = date
	res.Status = status

	return res, nil
}

// GetRecordById return a record where the provided ID appears
func GetRecordById(recordId int) (types.EventLog, error) {
	db, err := getDB()
	if err != nil {
		return types.EventLog{}, err
	}

	var res types.EventLog
	var id int
	var lamp, date string
	var status bool

	events := db.QueryRow("SELECT * FROM event_logs WHERE id=? ORDER BY id DESC", recordId)
	if err = events.Scan(&id, &lamp, &date, &status); err != nil {
		return types.EventLog{}, err
	}

	res.Id = id
	res.Lamp = lamp
	res.Date = date
	res.Status = status

	return res, nil
}

// GetRecordByLamp return a record where the provided Lamp name
func GetRecordByLamp(recordLamp string) ([]types.EventLog, error) {
	db, err := getDB()
	if err != nil {
		return []types.EventLog{}, err
	}

	events, err := db.Query("SELECT * FROM event_logs WHERE lamp=? ORDER BY date DESC LIMIT 1", recordLamp)
	if err != nil {
		return []types.EventLog{}, err
	}

	var tmp types.EventLog
	var res []types.EventLog

	for events.Next() {
		var id int
		var lamp, date string
		var status bool
		err = events.Scan(&id, &lamp, &date, &status)
		if err != nil {
			return []types.EventLog{}, err
		}
		tmp.Id = id
		tmp.Lamp = lamp
		tmp.Date = date
		tmp.Status = status
		res = append(res, tmp)
	}

	return res, nil
}

// GetRecordsByDate return events between specified dates
func GetRecordsByDate(startDate, endDate string) ([]types.EventLog, error) {
	// TODO add limit
	db, err := getDB()
	if err != nil {
		return []types.EventLog{}, err
	}

	events, err := db.Query("SELECT * FROM event_logs WHERE date BETWEEN ? AND ?", startDate, endDate)
	if err != nil {
		return []types.EventLog{}, err
	}

	var tmp types.EventLog
	var res []types.EventLog

	for events.Next() {
		var id int
		var lamp, date string
		var status bool

		if err := events.Scan(&id, &lamp, &date, &status); err != nil {
			return []types.EventLog{}, err
		}
		tmp.Id = id
		tmp.Lamp = lamp
		tmp.Date = date
		tmp.Status = status
		res = append(res, tmp)
	}

	return res, nil
}

// GetAll the latest status of every lamp
func GetAll() ([]types.EventLog, error) {
	// TODO add limit
	db, err := getDB()
	if err != nil {
		return []types.EventLog{}, err
	}

	// Must check this when it's not past midnight
	events, err := db.Query(`SELECT *
	FROM (
    	SELECT *
    	FROM event_logs
    	ORDER BY date DESC
	) AS sorted_logs
	GROUP BY lamp;`)
	if err != nil {
		return []types.EventLog{}, err
	}

	var tmp types.EventLog
	var res []types.EventLog

	for events.Next() {
		var id int
		var lamp, date string
		var status bool
		err = events.Scan(&id, &lamp, &date, &status)
		if err != nil {
			return []types.EventLog{}, err
		}
		tmp.Id = id
		tmp.Lamp = lamp
		tmp.Date = date
		tmp.Status = status
		res = append(res, tmp)
	}

	return res, nil
}
