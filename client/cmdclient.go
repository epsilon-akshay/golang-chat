package main

import (
	"bufio"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

type Message struct {
	Name    string
	Message string
}

func main() {

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:4001", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	name := "akshay"
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		msg := Message{
			Name:    name,
			Message: text,
		}
		err := c.WriteJSON(msg)
		if err != nil {
			log.Println("write:", err)
			return
		}
	}
}
