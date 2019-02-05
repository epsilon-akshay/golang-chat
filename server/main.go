package main

import (
	"chat-golang/server/handler"
	"chat-golang/server/logger"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	port := 4001
	host := "localhost"
	logLevel := "info"
	logFormatType := "json"

	var clients map[*websocket.Conn]bool

	log := logger.New(os.Stdout, logFormatType, logLevel)

	router := handler.NewRouter(log, clients)

	log.Infof("litsening to %v:%v", host, port)
	http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), router)
}
