package db

import (
	"database/sql"

	"backend/types"
	"backend/vars"

	_ "github.com/go-sql-driver/mysql"
)

// getDB returns a database handler
func getDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", vars.ConnectionString)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(1000)
	return db, err
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
	defer db.Close()

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}
