package handler

import (
	"chat-golang/server/model"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func NewRouter(clients map[*websocket.Conn][]string, chatMessages chan model.Message, db *sql.DB) http.Handler {
	router := mux.NewRouter()
	//handleConn := handleBroadcastConn(clients, chatMessages, db)
	handleIsolConn := handleIsolatedConn(clients, chatMessages, db)
	//router.HandleFunc("/", handleConn)
	router.HandleFunc("/isol", handleIsolConn)
	return router
}
