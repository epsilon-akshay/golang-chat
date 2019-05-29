package handler

import (
	"chat-golang/server/database"
	"chat-golang/server/logger"
	"chat-golang/server/messagehandler"
	"chat-golang/server/model"
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

func handleBroadcastConn(clients map[*websocket.Conn]bool, chatMessages chan model.Message, db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Log.Fatalf("could not upgrade http to websocket %v", err)
		}
		defer conn.Close()

		var msg model.Message
		rows, err := database.ReadMessageFromDb(db)
		if err != nil {
			logger.Log.Errorf("error in rows :%v", err)
		}

		for rows.Next() {
			err = rows.Scan(&msg.Time, &msg.SourceName, &msg.Message, &msg.Chatgroup)
			if err != nil {
				logger.Log.Error(err)
			}
			fmt.Printf("%v: %v", msg.SourceName, msg.Message)
		}

		clients[conn] = true

		mutex := &sync.Mutex{}

		go messagehandler.HandleBroadCastMessage(clients, chatMessages)

		messagehandler.Handle(clients, conn, chatMessages, db, mutex)
	})
}
