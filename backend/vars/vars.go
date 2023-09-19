package vars

import (
	"fmt"
	"os"
)

var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		"smarthome")
)

func GetPort() string {
	var port = os.Getenv("SH_PORT")
	if port == "" {
		port = "8088"
	}
	return port
}
