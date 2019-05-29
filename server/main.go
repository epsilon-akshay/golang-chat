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
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := 4001
	host := "0.0.0.0"
	logLevel := "info"
	logFormatType := "json"
	chatdb := "./chatdb.db"

	clients := make(map[*websocket.Conn][]string)
	chatMessages := make(chan model.Message)

	logger.New(os.Stdout, logFormatType, logLevel)

	db, err := sql.Open("sqlite3", chatdb)
	if err != nil {
		logger.Log.Fatal(err)
	}

	router := handler.NewRouter(clients, chatMessages, db)

	logger.Log.Infof("litsening to %v:%v", host, port)
	http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), router)
}
