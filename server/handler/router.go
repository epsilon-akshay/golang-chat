package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func NewRouter(log *logrus.Logger, clients map[*websocket.Conn]bool) http.Handler {
	router := mux.NewRouter()
	handleConn := handleConn(log, clients)
	router.HandleFunc("/", handleConn)
	return router
}
