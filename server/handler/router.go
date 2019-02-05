package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", handleConn)
	return router
}
