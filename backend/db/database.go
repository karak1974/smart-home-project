package db

import (
	"database/sql"
	"smarthome/types"
	"smarthome/vars"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() (*sql.DB, error) {
	return sql.Open("mysql", vars.ConnectionString)
}

// addRecord adds a log record about state of a lamp
// We don't give the ID due the database will create it
func AddRecord(eventLog types.EventLog) error {
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

// getRecordById return a record where the provided ID appears
func GetRecordById(recordId int) (types.EventLog, error) {
	db, err := getDB()
	if err != nil {
		return types.EventLog{}, err
	}

	events, err := db.Query("SELECT * FROM event_logs WHERE id=? LIMIT 1", recordId)
	if err != nil {
		return types.EventLog{}, err
	}

	var res types.EventLog
	var id int
	var lamp, date string
	var status bool

	if err = events.Scan(&id, &lamp, &date, &status); err != nil {
		return types.EventLog{}, err
	}

	res.Id = id
	res.Lamp = lamp
	res.Date = date
	res.Status = status

	return res, nil
}

// getRecordByLamp return a record where the provided Lamp name
func GetRecordByLamp(recordLamp string) ([]types.EventLog, error) {
	db, err := getDB()
	if err != nil {
		return []types.EventLog{}, err
	}

	events, err := db.Query("SELECT * FROM event_logs WHERE lamp=?", recordLamp)
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

// getRecordByDate return events between specified dates
func GetRecordByDate(startDate, endDate string) ([]types.EventLog, error) {
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

// getAll the latest state of every lamp
func GetAll() ([]types.EventLog, error) {
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
