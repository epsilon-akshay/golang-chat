package main

import (
	"chat-golang/server/handler"
	"chat-golang/server/logger"
	"chat-golang/server/model"
	"database/sql"
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
	chatdb := "./chatdb.db"

	clients := make(map[*websocket.Conn]bool)
	chatMessages := make(chan model.Message)

	log := logger.New(os.Stdout, logFormatType, logLevel)

	db, err := sql.Open("sqlite3", chatdb)
	if err != nil {
		log.Fatal(err)
	}

	router := handler.NewRouter(log, clients, chatMessages, db)

	log.Infof("litsening to %v:%v", host, port)
	http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), router)
}
