package clienthandler

import (
	"chat-golang/server/model"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func HandleMessage(clients map[*websocket.Conn]bool, chatMessages chan model.Message, log *logrus.Logger) error {
	for {
		msg := <-chatMessages

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Errorf("message marshal error %v", err)
				client.Close()
			}
		}
	}
}
