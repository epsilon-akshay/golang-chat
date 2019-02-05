package handler

import (
	"chat-golang/server/clienthandler"
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

		go clienthandler.HandleMessage(clients, chatMessages, log)
		clienthandler.Handle(clients, conn, chatMessages)

	})
}
