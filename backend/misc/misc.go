package misc

import (
	"backend/db"
	"backend/types"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
)

// FileServer serve a static file server based on the given folder
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, ":*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}

// SetupLamps read ROOMS system environment and assign one lamp to it
// This function os called only once before everything else starts
func SetupLamps() error {
	rooms := os.Getenv("ROOMS")
	lamps := strings.Fields(rooms)

	slog.Info("Adding lamps to the database", slog.Any("lamps", lamps))
	for _, lamp := range lamps {
		l := types.Lamp{
			Lamp:   lamp,
			Status: false,
		}
		if _, dbErr := db.AddRecord(l); dbErr != nil {
			slog.Error("Error adding lamp to the database", slog.Any("error", dbErr),
				slog.String("lamp", lamp))
			return dbErr
		}
	}
	return nil
}

// GetLamps return the array of lamps
func GetLamps() ([]string, error) {
	lamps, err := db.GetDistinctLamp()
	if err != nil {
		return nil, err
	}
	return lamps, nil
}
