package middleware

import (
	"fmt"
	"net/http"

	"bh/do.it/models"
	"github.com/gorilla/sessions"
)

var skipAuthRoutes = map[string]bool{
	"/":               true,
	"/GoogleLogin":    true,
	"/GoogleCallback": true,
}

// AuthHandler - middleware to protect all routes not in allowedPaths
type AuthHandler struct {
	handler     http.Handler
	sessionName string
	store       *sessions.CookieStore
}

// NewAuthHandler - return an instance of our AuthHandler
func NewAuthHandler(handler http.Handler, sessionName string, store *sessions.CookieStore) *AuthHandler {
	return &AuthHandler{handler: handler, sessionName: sessionName, store: store}
}

func (a *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	// skip authentication because of route (login)
	if skipAuthRoutes[path] == true {
		a.handler.ServeHTTP(w, r)
		return
	}

	session, err := a.store.Get(r, a.sessionName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// user must be logged in to see these routes
	user := session.Values["user"]
	if _, ok := user.(*models.User); ok {
		a.handler.ServeHTTP(w, r)
		return
	}

	// unauthorized. user is not logged in.
	w.WriteHeader(403)
}
