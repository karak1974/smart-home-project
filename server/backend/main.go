package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"backend/db"
	"backend/misc"
	"backend/router"
	"backend/vars"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	// Wait for MySQL to start
	slog.Info("Starting Smarthome API")
	for i := 0; i < vars.GetMaxTry(); i++ {
		if err := db.HealthCheck(); err == nil {
			slog.Info("Database reached", slog.Int("attempt", i+1))
			break
		} else {
			slog.Info("Tried to reaching database", slog.Int("attempt", i+1))
		}

		if i == vars.GetMaxTry()-1 {
			slog.Error("Could not connect to the database, exiting...")
			os.Exit(1)
		}

		time.Sleep(10 * time.Second)
	}

	// Setup lamps
	if err := misc.SetupLamps(); err != nil {
		slog.Error("Error during lamp setup", slog.Any("error", err))
		os.Exit(1)
	}

	// HTTP
	go func() {
		var port = vars.GetPort()
		r := chi.NewRouter()
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		})
		// Backend
		r.Use(c.Handler)
		r.Post("/api/addRecord", router.AddRecordHandler)
		r.Get("/api/getLastByLamp/{lamp}", router.GetLastByLampHandler)
		r.Get("/api/getLamps", router.GetLamps)
		r.Get("/api/hc", router.HealthCheckHandler)
		// Frontend
		router.FileServer(r, "/", http.Dir("./frontend"))

		slog.Info("Smarthome API is running", slog.String("port", port))
		if err := http.ListenAndServe(":"+port, r); err != nil {
			slog.Error("Could not serve HTTP API", slog.String("port", port))
		}
	}()

	// WS
	go func() {
		var port = vars.GetWSPort()

		slog.Info("Smarthome WS is running", slog.String("port", port), slog.String("path", "/smart-home"))

		http.HandleFunc("/smart-home", router.HandleClient)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			slog.Error("Error serving WebSocket port", slog.String("port", port))
		}
	}()

	select {}
}
