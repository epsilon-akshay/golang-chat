package main

import (
	"chat-golang/server/logger"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := 4001
	host := "localhost"
	logLevel := "info"
	logFormatType := "json"

	log := logger.New(os.Stdout, logFormatType, logLevel)

	log.Infof("litsening to %v:%v", host, port)
	http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), nil)
}
