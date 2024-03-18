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
		slog.Info("Trying reaching the database", slog.Int("attempt", i+1))

		if err := db.HealthCheck(); err == nil {
			slog.Info("Database reached")
			break
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

	// Setup server
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
	r.Use(c.Handler)
	// Backend
	r.Post("/api/addRecord", router.AddRecordHandler)
	r.Get("/api/getLastByLamp/{lamp}", router.GetLastByLampHandler)
	r.Get("/api/getLamps", router.GetLamps)
	r.Get("/api/hc", router.HealthCheckHandler)
	// WS
	r.Get("/smart-home", router.HandleClient)
	// Frontend
	misc.FileServer(r, "/", http.Dir("./frontend"))

	// Run server
	var port = vars.GetPort()
	slog.Info("Smarthome API is running", slog.String("port", port))
	if err := http.ListenAndServe(":"+port, r); err != nil {
		slog.Error("Could not serve HTTP API", slog.String("port", port))
	}
}
