package main

import (
	"fmt"
	"log"
	"net/http"
)

// HandleError - log error, send http response if possible
func HandleError(err error, w http.ResponseWriter) {
	if w != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\":\"" + err.Error() + "\"}"))
	} else {
		log.Fatal(err)
	}
}
