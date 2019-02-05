package clienthandler

import "github.com/gorilla/websocket"

func Handle(clients map[*websocket.Conn]bool, conn *websocket.Conn) {
	clients[conn] = true
}
