package handler

import (
	"chat-golang/server/clienthandler"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

func handleConn(log *logrus.Logger, clients map[*websocket.Conn]bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatalf("could not upgrade http to websocket %v", err)
		}
		defer conn.Close()

		clienthandler.Handle(clients, conn)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}
