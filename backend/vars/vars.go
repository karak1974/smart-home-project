package vars

import (
	"fmt"
	"os"
)

var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		"mysql",
		//os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		"smarthome")
)

// var Port = os.Getenv("SH_PORT")

var Port = "8088"
