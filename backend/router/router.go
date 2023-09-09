package router

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"smarthome/types"

	"github.com/go-chi/chi"
)

func AddRecordHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Got AddRecord POST request")

	event := &types.EventLog{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Info("Request body",
		slog.String("lamp", event.Lamp),
		slog.Bool("status", event.Status))

	w.Write([]byte("OK"))
}

func GetRecordByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	slog.Info("Got GetRecordById request",
		slog.String("id", id))
	w.Write([]byte("OK"))
}

func GetRecordByLampHandler(w http.ResponseWriter, r *http.Request) {
	lamp := chi.URLParam(r, "lamp")
	slog.Info("Got GetRecordByLamp request",
		slog.String("lamp", lamp))
	w.Write([]byte("OK"))
}

func GetRecordByDateHandler(w http.ResponseWriter, r *http.Request) {
	start := chi.URLParam(r, "start")
	end := chi.URLParam(r, "end")
	slog.Info("Got GetRecordByDate request",
		slog.String("start", start),
		slog.String("end", end))
	w.Write([]byte("OK"))
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Got GetAll request")
	w.Write([]byte("OK"))
}
