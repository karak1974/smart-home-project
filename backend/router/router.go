package router

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"smarthome/db"
	"smarthome/types"
	"strconv"

	"github.com/go-chi/chi"
)

func AddRecordHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Got AddRecord POST request")

	event := &types.EventLog{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}
	slog.Info("Request body",
		slog.String("lamp", event.Lamp),
		slog.Bool("status", event.Status))

	// Create and add a record to the database
	record := types.EventLog{
		Id:     0,
		Lamp:   event.Lamp,
		Date:   "",
		Status: event.Status,
	}
	record, err := db.AddRecord(record)
	if err != nil {
		slog.Error("Error adding record to the database",
			slog.String("error", err.Error()),
			slog.Any("record", record))
		http.Error(w, "Error adding record to the database", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(record)
	if err != nil {
		slog.Error("Error marshalling response")
	}

	if _, err = w.Write(resp); err != nil {
		slog.Error("Cound not serve request for AddRecord")
	}
}

func GetRecordByIdHandler(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "id")
	slog.Info("Got GetRecordById request",
		slog.String("id", urlId))

	id, err := strconv.Atoi(urlId)
	if err != nil {
		slog.Error("Error parsing request",
			slog.String("error", err.Error()),
			slog.String("record_id", urlId))
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	record, err := db.GetRecordById(id)
	if err != nil {
		slog.Error("Error getting record from the database",
			slog.String("error", err.Error()),
			slog.Int("record_id", id))
		http.Error(w, "Error getting record from the database", http.StatusInternalServerError)
	}

	resp, err := json.Marshal(record)
	if err != nil {
		slog.Error("Error marshalling response")
	}

	if _, err = w.Write(resp); err != nil {
		slog.Error("Cound not serve request for GetRecordById")
	}
}

func GetRecordByLampHandler(w http.ResponseWriter, r *http.Request) {
	lamp := chi.URLParam(r, "lamp")
	slog.Info("Got GetRecordByLamp request",
		slog.String("lamp", lamp))

	record, err := db.GetRecordByLamp(lamp)
	if err != nil {
		slog.Error("Error getting record from the database",
			slog.String("error", err.Error()),
			slog.String("lamp", lamp))
		http.Error(w, "Error getting record from the database", http.StatusInternalServerError)
	}

	resp, err := json.Marshal(record)
	if err != nil {
		slog.Error("Error marshalling response")
	}

	if _, err = w.Write(resp); err != nil {
		slog.Error("Cound not serve request for GetRecordByLamp")
	}
}

func GetRecordsByDateHandler(w http.ResponseWriter, r *http.Request) {
	start := chi.URLParam(r, "start")
	end := chi.URLParam(r, "end")
	slog.Info("Got GetRecordByDate request",
		slog.String("start", start),
		slog.String("end", end))
	_, _ = w.Write([]byte("OK"))
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Got GetAll request")
	_, _ = w.Write([]byte("OK"))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("HealthCheck")
	// TODO add database connection check
	_, _ = w.Write([]byte("OK"))
}
