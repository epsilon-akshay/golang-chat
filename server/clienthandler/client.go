package clienthandler

import (
	"chat-golang/server/model"
	"fmt"

	"github.com/gorilla/websocket"
)

func Handle(clients map[*websocket.Conn]bool, conn *websocket.Conn, chatMessages chan model.Message) error {
	for {
		var msg model.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			delete(clients, conn)
			return fmt.Errorf("could not read message from client: %v", err)
		}

		chatMessages <- msg
	}
}
