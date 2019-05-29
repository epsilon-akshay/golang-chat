package messagehandler

import (
	"chat-golang/server/logger"
	"chat-golang/server/model"
	"fmt"

	"github.com/gorilla/websocket"
)

func HandleBroadCastMessage(clients map[*websocket.Conn]bool, chatMessages chan model.Message) error {
	for {
		msg := <-chatMessages

		for client := range clients {
			err := client.WriteJSON(msg)
			fmt.Println(msg)
			if err != nil {
				logger.Log.Errorf("message marshal error %v", err)
				client.Close()
			}
		}
	}
}
