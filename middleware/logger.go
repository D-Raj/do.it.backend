package middleware

import (
	"fmt"
	"net/http"
)

// LoggerHandler - middleware to log info about each request
type LoggerHandler struct {
	handler http.Handler
}

// NewLoggerHandler - return an instance of our LoggerHandler
func NewLoggerHandler(handler http.Handler) http.Handler {
	return &LoggerHandler{handler: handler}
}

func (h *LoggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	path := r.URL.Path
	message := method + ": " + path

	fmt.Println(message)

	h.handler.ServeHTTP(w, r)
}
