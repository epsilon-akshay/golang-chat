package handler

import (
	"chat-golang/server/database"
	"chat-golang/server/messagehandler"
	"chat-golang/server/model"
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

func handleConn(log *logrus.Logger, clients map[*websocket.Conn]bool, chatMessages chan model.Message, db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatalf("could not upgrade http to websocket %v", err)
		}
		defer conn.Close()

		var msg model.Message
		rows, err := database.ReadMessageFromDb(db)
		if err != nil {
			log.Errorf("error in rows :%v", err)
		}

		for rows.Next() {
			err = rows.Scan(&msg.Time, &msg.Name, &msg.Message, &msg.Chatgroup)
			if err != nil {
				log.Error(err)
			}
			fmt.Printf("%v: %v", msg.Name, msg.Message)
		}

		clients[conn] = true

		mutex := &sync.Mutex{}

		go messagehandler.HandleMessage(clients, chatMessages, log)

		messagehandler.Handle(clients, conn, chatMessages, db, mutex)
	})
}
