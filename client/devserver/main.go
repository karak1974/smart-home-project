package main

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleClient(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error connection to upgrade", slog.Any("error", err))
		return
	}
	defer conn.Close()

	// Send "ACK" to the client
	if err := conn.WriteMessage(websocket.TextMessage, []byte("ACK")); err != nil {
		slog.Error("Error write message", slog.Any("error", err))
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Error read message", slog.Any("error", err))
			return
		}
		slog.Info("Received message from client", slog.String("message", string(p)))

		if err := conn.WriteMessage(messageType, p); err != nil {
			slog.Error("Error write", slog.Any("error", err))
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
