package main

import (
	"log/slog"
	"net/http"
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
	r.Get("/getLast", router.GetLastHandler)
	r.Get("/getLast/{amount}", router.GetLastAmountHandler)
	r.Get("/hc", router.HealthCheckHandler)

	slog.Info("API starting",
		slog.String("port", vars.Port))

	if err := http.ListenAndServe(":"+vars.Port, r); err != nil {
		slog.Error("Cound not serve HTTP API",
			slog.String("port", vars.Port))
	}
}
