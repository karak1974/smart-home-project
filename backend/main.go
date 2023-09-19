package main

import (
	"log/slog"
	"net/http"
	"os"
	"smarthome/router"
	"smarthome/vars"

	"github.com/go-chi/chi"
)

func main() {
	slog.Info("Starting SmartHome backend")

	r := chi.NewRouter()

	// Define your handler functions
	r.Post("/addRecord", router.AddRecordHandler)
	r.Get("/getRecordById/{id}", router.GetRecordByIdHandler)
	r.Get("/getRecordByLamp/{lamp}", router.GetRecordByLampHandler)
	r.Get("/getRecordByDate/{start}/{end}", router.GetRecordsByDateHandler)
	r.Get("/getAll", router.GetAllHandler)
	//r.Get("/getLast", router.GetLastHandler)
	//r.Get("/getLast/{amount}", router.GetLastAmountHandler)
	r.Get("/hc", router.HealthCheckHandler)

	slog.Info("API starting",
		slog.String("port", vars.Port))
	slog.Info("DEBUG",
		slog.String("DB_USER", os.Getenv("DB_USER")),
		slog.String("DB_PASS", os.Getenv("DB_PASS")),
		slog.String("DB_PORT", os.Getenv("DB_PORT")))
	http.ListenAndServe(":"+vars.Port, r)
}
