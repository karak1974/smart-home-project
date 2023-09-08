package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"log/slog"
	"net/http"
)

func AddRecordHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Got AddRecord POST request")
	fmt.Fprintf(w, "Handling POST request to addRecord")
}

func GetRecordByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	slog.Info("Got GetRecordById request",
		slog.String("id", id))
	fmt.Fprintf(w, "Handling GET request for record with ID: %s", id)
}

func GetRecordByLampHandler(w http.ResponseWriter, r *http.Request) {
	lamp := chi.URLParam(r, "lamp")
	slog.Info("Got GetRecordByLamp request",
		slog.String("lamp", lamp))
	fmt.Fprintf(w, "Handling GET request for record by Lamp: %s", lamp)
}

func GetRecordByDateHandler(w http.ResponseWriter, r *http.Request) {
	start := chi.URLParam(r, "start")
	end := chi.URLParam(r, "end")
	slog.Info("Got GetRecordByDate request",
		slog.String("start", start),
		slog.String("end", end))
	fmt.Fprintf(w, "Handling GET request for records between %s and %s", start, end)
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Got GetAll request")
	fmt.Fprintf(w, "Handling GET request to getAll")
}
