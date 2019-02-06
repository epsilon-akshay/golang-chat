package handler

import (
	"chat-golang/server/messagehandler"
	"chat-golang/server/model"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

func handleConn(log *logrus.Logger, clients map[*websocket.Conn]bool, chatMessages chan model.Message) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatalf("could not upgrade http to websocket %v", err)
		}
		defer conn.Close()

		clients[conn] = true

		go messagehandler.HandleMessage(clients, chatMessages, log)

		messagehandler.Handle(clients, conn, chatMessages)
	})
}
