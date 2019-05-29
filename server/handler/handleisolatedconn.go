package handler

import (
	"chat-golang/server/logger"
	"chat-golang/server/model"
	"database/sql"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/websocket"
)

func handleIsolatedConn(clients map[*websocket.Conn][]string, chatMessages chan model.Message, db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		source := params["source_name"][0]    // set in db when registering
		dest := params["destination_name"][0] // set in db when registering

		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Log.Errorf("could not upgrade conn")
			return
		}
		defer conn.Close()

		clients[conn] = []string{source, dest}

		go func() {
			for {
				message := <-chatMessages
				destConn := []string{dest, source}
				for k, v := range clients { // think of a better way
					if reflect.DeepEqual(v, destConn) {
						err := k.WriteJSON(message)
						fmt.Println(message)
						if err != nil {
							logger.Log.Errorf("message marshal error %v", err)
						}

					}
				}
			}
		}()
		
		for {
			var msg model.Message
			err = conn.ReadJSON(&msg)
			if err != nil {
				logger.Log.Error("could not read message")
			}
			chatMessages <- msg
		}
	})
}
