package misc

import (
	"backend/db"
	"backend/types"
	"encoding/hex"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
)

var xorKey = func() []byte {
	keyString := os.Getenv("XOR_KEY")
	if keyString == "" {
		slog.Error("Environment variable not set", slog.String("env", "XOR_KEY"))
		os.Exit(1)
	}

	key, err := hex.DecodeString(keyString)
	if err != nil {
		slog.Error("Can't decode XOR key", slog.Any("error", err))
	}

	return key
}()

// XorData takes an input string and XOR it with a key
func XorData(input string) string {
	if len(input) != len(xorKey) {
		return ""
	}

	result := make([]byte, len(input))

	for i := range input {
		result[i] = input[i] ^ xorKey[i]
	}

	return string(result)
}

// FileServer serve a static file server based on the given folder
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
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
			Lamp:  lamp,
			State: false,
		}
		if _, dbErr := db.AddRecord(l); dbErr != nil {
			slog.Error("Error adding lamp to the database", slog.Any("error", dbErr),
				slog.String("lamp", lamp))
			return dbErr
		}
	}
	return nil
}
