package main

import (
	"encoding/hex"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var d = "11001010"

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

func xorData(input string) string {
	if len(input) != len(xorKey) {
		return ""
	}

	result := make([]byte, len(input))

	for i := range input {
		result[i] = input[i] ^ xorKey[i]
	}

	return string(result)
}

func handleClient(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error connection to upgrade", slog.Any("error", err))
		return
	}
	defer conn.Close()

	slog.Info("Controller connected")

	// SENDING DATA
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if strings.Contains(err.Error(), "connection reset by peer") {
				slog.Warn("Controller disconnected")
				return
			} else if strings.Contains(err.Error(), "connection timed out") {
				slog.Warn("Controloller timeouted")
				return
			} else {
				slog.Error("Error read message", slog.Any("error", err))
				return
			}
		}
		if string(msg) != "OK" {
			slog.Info("Received message from client", slog.String("client", conn.RemoteAddr().String()),
				slog.String("message", string(msg)))
		}

		// Send the XORed data
		if err := conn.WriteMessage(websocket.TextMessage, []byte(xorData(d))); err != nil {
			slog.Error("Error write message", slog.Any("error", err))
			return
		}
	}
}

func main() {
	slog.Info("DEV WebSocket server")
	slog.Info("Test connection on ws://127.0.0.1:8087/smart-home")

	http.HandleFunc("/smart-home", handleClient)
	if err := http.ListenAndServe(":8087", nil); err != nil {
		slog.Error("Error serving WebSocket port", slog.Int("port", 8087))
	}
}
