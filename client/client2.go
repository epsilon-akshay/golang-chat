package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Time            string
	SourceName      string
	DestinationName string
	Message         string
}

func main() {
	name := "pranav"
	destName := "akshay"
	c, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://localhost:4001/isol?source_name=%s&destination_name=%s", name, destName), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	go func() {
		var msg Message
		for {
			err := c.ReadJSON(&msg)
			if err != nil {
				fmt.Print("error reading message")
			}
			fmt.Println(msg)
		}
	}()
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		msg := Message{
			Time:            time.Now().String(),
			SourceName:      name,
			DestinationName: destName,
			Message:         text,
		}
		err := c.WriteJSON(msg)
		if err != nil {
			log.Println("write:", err)
			return
		}
	}
}
