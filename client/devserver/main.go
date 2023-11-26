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

	// SENDING DATA
	for {
		//messageType, msg, err := conn.ReadMessage()
		_, msg, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Error read message", slog.Any("error", err))
			return
		}
		if string(msg) != "OK" {
			slog.Info("Received message from client", slog.String("client", conn.RemoteAddr().String()),
				slog.String("message", string(msg)))
		}
		// Else do something if the controller won't response for 5 sec

		d := "11010011"
		if err := conn.WriteMessage(websocket.TextMessage, []byte(d)); err != nil {
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
