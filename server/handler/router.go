package handler

import (
	"chat-golang/server/model"
	"database/sql"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func NewRouter(log *logrus.Logger, clients map[*websocket.Conn]bool, chatMessages chan model.Message, db *sql.DB) http.Handler {
	router := mux.NewRouter()
	handleConn := handleConn(log, clients, chatMessages, db)
	router.HandleFunc("/", handleConn)
	return router
}
