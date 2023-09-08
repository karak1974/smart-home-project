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
	r.Get("/getRecordByDate/{start}/{end}", router.GetRecordByDateHandler)
	r.Get("/getAll", router.GetAllHandler)

	slog.Info("API starting",
		slog.String("port", vars.Port))
	http.ListenAndServe(":"+vars.Port, r)
}
