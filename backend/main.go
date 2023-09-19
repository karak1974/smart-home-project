package main

import (
	"log/slog"
	"net/http"
	"os"
	"smarthome/db"
	"time"

	"smarthome/router"
	"smarthome/vars"

	"github.com/go-chi/chi"
)

func main() {
	// Wait 30 sec for MySQL
	slog.Info("Waiting MySQL to start")
	time.Sleep(30 * time.Second)

	for i := 0; i < 20; i++ {
		slog.Info("Trying reaching the database",
			slog.Int("attempt", i+1))
		time.Sleep(5 * time.Second)
		if err := db.HealthCheck(); err == nil {
			break
		}

		if i == 20 {
			slog.Error("Could not connect to the database, exiting...")
			os.Exit(1)
		}
	}

	// Setup server
	r := chi.NewRouter()
	r.Post("/addRecord", router.AddRecordHandler)
	r.Get("/getRecordById/{id}", router.GetRecordByIdHandler)
	r.Get("/getRecordByLamp/{lamp}", router.GetRecordByLampHandler)
	r.Get("/getLast", router.GetLastHandler)
	r.Get("/getLast/{amount}", router.GetLastAmountHandler)
	r.Get("/hc", router.HealthCheckHandler)

	// Run server
	var port = vars.GetPort()
	slog.Info("Smarthome API starting",
		slog.String("port", port))
	if err := http.ListenAndServe(":"+port, r); err != nil {
		slog.Error("Could not serve HTTP API",
			slog.String("port", port))
	}
}
