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

// AddRecord adds a log record about status2 of a lamp
// We don't give the ID due the database will create it
func AddRecord(eventLog types.EventLog) (types.EventLog, error) {
	db, err := getDB()
	if err != nil {
		return types.EventLog{}, err
	}

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

	events := db.QueryRow("SELECT * FROM event_logs WHERE id=? ORDER BY id DESC LIMIT 1", recordId)
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
func GetRecordByLamp(recordLamp string) (types.EventLog, error) {
	db, err := getDB()
	if err != nil {
		return types.EventLog{}, err
	}

	events, err := db.Query("SELECT * FROM event_logs WHERE lamp=? ORDER BY date DESC LIMIT 1", recordLamp)
	if err != nil {
		return types.EventLog{}, err
	}

	var res types.EventLog

	if events.Next() {
		var id int
		var lamp, date string
		var status bool
		err = events.Scan(&id, &lamp, &date, &status)
		if err != nil {
			return types.EventLog{}, err
		}
		res.Id = id
		res.Lamp = lamp
		res.Date = date
		res.Status = status
	}

	return res, nil
}

// GetlastRecord return a record
func GetLastRecord() (types.EventLog, error) {
	db, err := getDB()
	if err != nil {
		return types.EventLog{}, err
	}

	events, err := db.Query("SELECT * FROM event_logs ORDER BY date DESC LIMIT 1")
	if err != nil {
		return types.EventLog{}, err
	}

	var res types.EventLog

	if events.Next() {
		var id int
		var lamp, date string
		var status bool
		err = events.Scan(&id, &lamp, &date, &status)
		if err != nil {
			return types.EventLog{}, err
		}
		res.Id = id
		res.Lamp = lamp
		res.Date = date
		res.Status = status
	}

	return res, nil
}

// GetLastAmountRecord returns records by given amount
func GetLastAmountRecord(amount int) ([]types.EventLog, error) {
	db, err := getDB()
	if err != nil {
		return []types.EventLog{}, err
	}

	stmt, err := db.Prepare("SELECT * FROM event_logs ORDER BY date DESC LIMIT ?")
	if err != nil {
		return []types.EventLog{}, err
	}
	defer stmt.Close()

	events, err := stmt.Query(amount)
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
