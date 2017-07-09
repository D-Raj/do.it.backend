package main

import (
	"fmt"
	"net/http"
	"os"
)

// StaticHandler - do stuff get fucked
func StaticHandler(w http.ResponseWriter, r *http.Request) {

	wwwRoot := os.Getenv("STATIC_DIR")
	fmt.Println(wwwRoot)
	// never serve static (etc) shit when not GET
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	// serve index.html at root, else look for file
	if r.URL.Path == "/" {
		session, err := store.Get(r, sessionName)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(500), 500)
		}
		user := session.Values["user"]
		if user != nil {
			http.Redirect(w, r, "/api/me", http.StatusTemporaryRedirect)
		}
		http.ServeFile(w, r, wwwRoot+"/index.html")
	}
	http.ServeFile(w, r, wwwRoot+r.URL.Path)
}
