package db

import (
	"database/sql"

	"backend/types"
	"backend/vars"

	_ "github.com/go-sql-driver/mysql"
)

// getDB returns a database handler
func getDB() (*sql.DB, error) {
	return sql.Open("mysql", vars.ConnectionString)
}

// AddRecord adds a log record about status2 of a lamp
// We don't give the ID due the database will create it
func AddRecord(eventLog types.Lamp) (types.Lamp, error) {
	db, err := getDB()
	if err != nil {
		return types.Lamp{}, err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO event_logs (lamp, date, status) VALUES (?, NOW(), ?)",
		eventLog.Lamp,
		eventLog.Status)
	if err != nil {
		return types.Lamp{}, err
	}

	var res types.Lamp
	var id int
	var lamp, date string
	var status bool

	events := db.QueryRow("SELECT * FROM event_logs WHERE lamp=? ORDER BY id DESC", eventLog.Lamp)
	if err = events.Scan(&id, &lamp, &date, &status); err != nil {
		return types.Lamp{}, err
	}

	res.Id = id
	res.Lamp = lamp
	res.Date = date
	res.Status = status

	return res, nil
}

// GetRecordById return a record where the provided ID appears
func GetRecordById(recordId int) (types.Lamp, error) {
	db, err := getDB()
	if err != nil {
		return types.Lamp{}, err
	}
	defer db.Close()

	var res types.Lamp
	var id int
	var lamp, date string
	var status bool

	events := db.QueryRow("SELECT * FROM event_logs WHERE id=? ORDER BY id DESC LIMIT 1", recordId)
	if err = events.Scan(&id, &lamp, &date, &status); err != nil {
		return types.Lamp{}, err
	}

	res.Id = id
	res.Lamp = lamp
	res.Date = date
	res.Status = status

	return res, nil
}

// GetLastByLamp return a record with the provided lamp's name
func GetLastByLamp(recordLamp string) (types.Lamp, error) {
	db, err := getDB()
	if err != nil {
		return types.Lamp{}, err
	}
	defer db.Close()

	events, err := db.Query("SELECT * FROM event_logs WHERE lamp=? ORDER BY date DESC LIMIT 1", recordLamp)
	if err != nil {
		return types.Lamp{}, err
	}

	var res types.Lamp

	if events.Next() {
		var id int
		var lamp, date string
		var status bool
		err = events.Scan(&id, &lamp, &date, &status)
		if err != nil {
			return types.Lamp{}, err
		}
		res.Id = id
		res.Lamp = lamp
		res.Date = date
		res.Status = status
	}

	return res, nil
}

// GetLastRecord return the last record
func GetLastRecord() (types.Lamp, error) {
	db, err := getDB()
	if err != nil {
		return types.Lamp{}, err
	}
	defer db.Close()

	events, err := db.Query("SELECT * FROM event_logs ORDER BY date DESC LIMIT 1")
	if err != nil {
		return types.Lamp{}, err
	}

	var res types.Lamp

	if events.Next() {
		var id int
		var lamp, date string
		var status bool
		err = events.Scan(&id, &lamp, &date, &status)
		if err != nil {
			return types.Lamp{}, err
		}
		res.Id = id
		res.Lamp = lamp
		res.Date = date
		res.Status = status
	}

	return res, nil
}

// GetLastAmountRecord returns last records by given amount
func GetLastAmountRecord(amount int) ([]types.Lamp, error) {
	db, err := getDB()
	if err != nil {
		return []types.Lamp{}, err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM event_logs ORDER BY date DESC LIMIT ?")
	if err != nil {
		return []types.Lamp{}, err
	}
	defer stmt.Close()

	events, err := stmt.Query(amount)
	if err != nil {
		return []types.Lamp{}, err
	}

	var tmp types.Lamp
	var res []types.Lamp

	for events.Next() {
		var id int
		var lamp, date string
		var status bool
		err = events.Scan(&id, &lamp, &date, &status)
		if err != nil {
			return []types.Lamp{}, err
		}
		tmp.Id = id
		tmp.Lamp = lamp
		tmp.Date = date
		tmp.Status = status
		res = append(res, tmp)
	}

	return res, nil
}

// GetLastAmountByLamp returns last records for a specific lamps by given amount
func GetLastAmountByLamp(reqLamp string, amount int) ([]types.Lamp, error) {
	db, err := getDB()
	if err != nil {
		return []types.Lamp{}, err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM event_logs WHERE lamp=? ORDER BY date DESC LIMIT ?")
	if err != nil {
		return []types.Lamp{}, err
	}
	defer stmt.Close()

	events, err := stmt.Query(reqLamp, amount)
	if err != nil {
		return []types.Lamp{}, err
	}

	var tmp types.Lamp
	var res []types.Lamp

	for events.Next() {
		var id int
		var lamp, date string
		var status bool
		err = events.Scan(&id, &lamp, &date, &status)
		if err != nil {
			return []types.Lamp{}, err
		}
		tmp.Id = id
		tmp.Lamp = lamp
		tmp.Date = date
		tmp.Status = status
		res = append(res, tmp)
	}

	return res, nil
}

func GetDistinctLamp() ([]string, error) {
	db, err := getDB()
	if err != nil {
		return []string{}, err
	}
	defer db.Close()

	var res []string
	lampArray, err := db.Query("SELECT DISTINCT lamp FROM event_logs;")
	if err != nil {
		return []string{}, err
	}
	defer lampArray.Close()

	for lampArray.Next() {
		var lamp string
		err = lampArray.Scan(&lamp)
		if err != nil {
			return []string{}, err
		}
		res = append(res, lamp)
	}

	return res, nil
}

func HealthCheck() error {
	db, err := getDB()
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}
