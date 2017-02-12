package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"bh/do.it/middleware"
	"bh/do.it/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

var db *sqlx.DB
var store = sessions.NewCookieStore([]byte("this is the hook. it's catchy. you like it."))
var sessionName = "i-like-the-grooves.but-i-digress."

func init() {
	gob.Register(&models.User{})
}

func main() {
	models.InitDB("do_it_dev:do_it_dev@/do_it_dev")

	router := httprouter.New()

	/* authentication */
	router.GET("/", DebugIndex) // TEMP for dev/debug
	router.GET("/GoogleLogin", HandleGoogleLogin)
	router.GET("/GoogleCallback", HandleGoogleCallback)
	router.GET("/logout", LogoutHandler)

	/* actions read/write */
	router.GET("/me", AllActionsHandler)
	router.POST("/me", NewActionHandler)

	/* goals read/write */
	router.GET("/me/goals", AllGoalsHandler)
	router.POST("/me/goals", NewGoalHandler)

	authHandler := middleware.NewAuthHandler(router, sessionName, store)
	loggerHandler := middleware.NewLoggerHandler(authHandler)

	fmt.Println("Server listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", context.ClearHandler(loggerHandler)))
}
