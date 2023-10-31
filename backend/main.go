package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"smarthome/db"
	"smarthome/router"
	"smarthome/vars"

	"github.com/go-chi/chi"
)

func main() {
	// Wait for MySQL to start
	slog.Info("Starting Smarthome API")
	for i := 0; i < vars.GetMaxTry(); i++ {
		slog.Info("Trying reaching the database",
			slog.Int("attempt", i+1))
		time.Sleep(5 * time.Second)
		if err := db.HealthCheck(); err == nil {
			break
		}

		if i == vars.GetMaxTry()-1 {
			slog.Error("Could not connect to the database, exiting...")
			os.Exit(1)
		}
	}

	// Setup server
	r := chi.NewRouter()
	r.Post("/addRecord", router.AddRecordHandler)
	r.Get("/getRecordById/{id}", router.GetRecordByIdHandler)
	r.Get("/getLastByLamp/{lamp}", router.GetLastByLampHandler)
	r.Get("/getLast", router.GetLastHandler)
	r.Get("/getLast/{amount}", router.GetLastAmountHandler)
	r.Get("/getLast/{lamp}/{amount}", router.GetLastAmountByLampHandler)
	r.Get("/hc", router.HealthCheckHandler)

	// Run server
	var port = vars.GetPort()
	slog.Info("Smarthome API is running",
		slog.String("port", port))
	if err := http.ListenAndServe(":"+port, r); err != nil {
		slog.Error("Could not serve HTTP API",
			slog.String("port", port))
	}
}
