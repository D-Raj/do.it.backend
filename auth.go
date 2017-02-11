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
)

// DebugIndex - route for login while working on backend. Remove when frontend is built
func DebugIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	htmlIndex := `<html><body><a href="/GoogleLogin">Log in with Google</a></body></html>`
	fmt.Fprintf(w, htmlIndex)
}

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

	var user map[string]interface{}
	if err := json.Unmarshal(contents, &user); err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
	}

	id, _ := user["id"].(string)
	name, _ := user["name"].(string)
	email, _ := user["email"].(string)

	_, err = models.AddUser("GOOGLE", id, name, email)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
	}

	fmt.Fprintf(w, "Logged in, hooray")
}
