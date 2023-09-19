package router

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"smarthome/db"
	"smarthome/types"

	"github.com/go-chi/chi"
)

// AddRecordHandler handler for /addRecord POST request
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
			slog.Any("error", err),
			slog.Any("record", record))
		http.Error(w, "Error adding record to the database", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(record)
	if err != nil {
		slog.Error("Error marshalling response")
	}

	if _, err = w.Write(resp); err != nil {
		slog.Error("Could not serve request for AddRecord")
	}
}

// GetRecordByIdHandler handler for /getRecordById/<ID> GET request
func GetRecordByIdHandler(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "id")
	slog.Info("Got GetRecordById GET request",
		slog.String("id", urlId))

	id, err := strconv.Atoi(urlId)
	if err != nil {
		slog.Error("Error parsing request",
			slog.Any("error", err),
			slog.String("record_id", urlId))
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	record, err := db.GetRecordById(id)
	if err != nil {
		slog.Error("Error getting record from the database",
			slog.Any("error", err),
			slog.Int("record_id", id))
		http.Error(w, "Error getting record from the database", http.StatusInternalServerError)
	}

	resp, err := json.Marshal(record)
	if err != nil {
		slog.Error("Error marshalling response")
	}

	if _, err = w.Write(resp); err != nil {
		slog.Error("Could not serve request for GetRecordById")
	}
}

// GetRecordByLampHandler handler for /getRecordByLamp/<LAMP> GET requests
func GetRecordByLampHandler(w http.ResponseWriter, r *http.Request) {
	lamp := chi.URLParam(r, "lamp")
	slog.Info("Got GetRecordByLamp GET request",
		slog.String("lamp", lamp))

	record, err := db.GetRecordByLamp(lamp)
	if err != nil {
		slog.Error("Error getting record from the database",
			slog.Any("error", err),
			slog.String("lamp", lamp))
		http.Error(w, "Error getting record from the database", http.StatusInternalServerError)
	}

	resp, err := json.Marshal(record)
	if err != nil {
		slog.Error("Error marshalling response")
	}

	if _, err = w.Write(resp); err != nil {
		slog.Error("Could not serve request for GetRecordByLamp")
	}
}

// GetLastHandler handler for /getLast GET requests
func GetLastHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Got GetLast request")

	record, err := db.GetLastRecord()
	if err != nil {
		slog.Error("Error getting record from the database",
			slog.Any("error", err))
		http.Error(w, "Error getting record from the database", http.StatusInternalServerError)
	}

	resp, err := json.Marshal(record)
	if err != nil {
		slog.Error("Error marshalling response")
	}

	if _, err = w.Write(resp); err != nil {
		slog.Error("Could not serve request for GetLast")
	}
}

// GetLastAmountHandler handler for /getLast/<AMOUNT> GET requests
func GetLastAmountHandler(w http.ResponseWriter, r *http.Request) {
	urlAmount := chi.URLParam(r, "amount")
	slog.Info("Got GetLastAmount GET request",
		slog.String("amount", urlAmount))

	amount, err := strconv.Atoi(urlAmount)
	if err != nil {
		slog.Error("Error parsing request",
			slog.Any("error", err),
			slog.String("amount", urlAmount))
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	record, err := db.GetLastAmountRecord(amount)
	if err != nil {
		slog.Error("Error getting record from the database",
			slog.Any("error", err))
		http.Error(w, "Error getting record from the database", http.StatusInternalServerError)
	}

	resp, err := json.Marshal(record)
	if err != nil {
		slog.Error("Error marshalling response")
	}

	if _, err = w.Write(resp); err != nil {
		slog.Error("Could not serve request for GetLast")
	}
}

// HealthCheckHandler handler for /hc GET requests
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("HealthCheck")

	var resp = "OK"
	if err := db.HealthCheck(); err != nil {
		resp = "NOT_OK"
		slog.Error("Could not connect to the database")
	}
	if _, err := w.Write([]byte(resp)); err != nil {
		slog.Error("Could not serve request for HealthCheck",
			slog.Any("error", err))
	}
}
