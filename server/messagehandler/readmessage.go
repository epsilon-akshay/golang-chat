package messagehandler

import (
	"chat-golang/server/database"
	"chat-golang/server/model"
	"database/sql"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

func Handle(clients map[*websocket.Conn]bool, conn *websocket.Conn, chatMessages chan model.Message, db *sql.DB, mutex *sync.Mutex) error {
	for {
		var msg model.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			delete(clients, conn)
			return fmt.Errorf("could not read message from client: %v", err)
		}

		mutex.Lock()
		database.WriteToDb(db)
		mutex.Unlock()

		chatMessages <- msg
	}
}
