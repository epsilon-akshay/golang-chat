package model

import (
	"github.com/gorilla/websocket"
)

type ClientInfo struct {
	Name string
	Conn *websocket.Conn
}
