package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"bh/do.it/models"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// TODO generate this random for each user
	oauthStateString = "random"

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/GoogleCallback",
		ClientID:     os.Getenv("GOOGLE_KEY"),
		ClientSecret: os.Getenv("GOOGLE_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	skipAuthRoutes = map[string]bool{
		"/":               true,
		"/GoogleLogin":    true,
		"/GoogleCallback": true,
	}
)

/** ******************************************************************** /
  Dev/Debug bullshit. Remove when possible
 ********************************************************************* **/

// DebugIndex - route for login while working on backend. Remove when frontend is built
func DebugIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// session this shit
	session, err := store.Get(r, sessionName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
	}
	user := session.Values["user"]
	if user != nil {
		http.Redirect(w, r, "/me", http.StatusTemporaryRedirect)
	}

	htmlIndex := `<html><body><a href="/GoogleLogin">Log in with Google</a></body></html>`
	fmt.Fprintf(w, htmlIndex)
}

/** ******************************************************************** /
  Google Handlers
 ********************************************************************* **/

// HandleGoogleLogin - send user to login with google
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleGoogleCallback - manage response from google when user tries to login
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("Invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	var rawUser map[string]interface{}
	if err := json.Unmarshal(contents, &rawUser); err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
	}

	externalID, _ := rawUser["id"].(string)
	name, _ := rawUser["name"].(string)
	email, _ := rawUser["email"].(string)

	internalID, err := models.GetUserID("GOOGLE", externalID, name, email)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
	}

	user := models.User{ID: internalID, Name: name, Email: email}

	// session this shit
	session, err := store.Get(r, sessionName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
	}
	session.Values["user"] = user
	session.Save(r, w)

	http.Redirect(w, r, "/me", http.StatusTemporaryRedirect)
}

/** ******************************************************************** /
  Authentication Middleware
 ********************************************************************* **/

// AuthHandler - middleware to protect all routes not in allowedPaths
type AuthHandler struct {
	handler http.Handler
}

// NewAuthHandler - return an instance of our AuthHandler
func NewAuthHandler(handler http.Handler) *AuthHandler {
	return &AuthHandler{handler: handler}
}

func (a *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	// skip authentication because of route (login)
	if skipAuthRoutes[path] == true {
		a.handler.ServeHTTP(w, r)
		return
	}

	session, err := store.Get(r, sessionName)
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

/** ******************************************************************** /
  Authentication Middleware
 ********************************************************************* **/

// LogoutHandler - log user out, destroy the session
func LogoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// session this shit
	session, err := store.Get(r, sessionName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
	}

	session.Values["user"] = nil
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
