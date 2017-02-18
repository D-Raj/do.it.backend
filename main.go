package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

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
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	models.InitDB("do_it_dev:do_it_dev@/do_it_dev")

	apiRouter := httprouter.New()
	authRouter := httprouter.New()

	/* authentication */
	authRouter.GET("/auth/GoogleLogin", HandleGoogleLogin)
	authRouter.GET("/auth/GoogleCallback", HandleGoogleCallback)
	authRouter.GET("/auth/logout", LogoutHandler)

	/* actions read/write */
	apiRouter.GET("/api/me", AllActionsHandler)
	apiRouter.POST("/api/me", NewActionHandler)

	/* goals read/write */
	apiRouter.GET("/api/me/goals", AllGoalsHandler)
	apiRouter.POST("/api/me/goals", NewGoalHandler)

	loggerHandler := middleware.NewLoggerHandler(authRouter)
	setHeadersHandler := middleware.NewSetHeadersHandler(loggerHandler)
	authHandler := middleware.NewAuthHandler(setHeadersHandler, sessionName, store)

	apiHandler := apiRouter // uneccesary for performance. reassign for pretty read

	mux := http.NewServeMux()
	mux.Handle("/auth/", authHandler)
	mux.Handle("/api/", apiHandler)
	mux.HandleFunc("/", StaticHandler)

	fmt.Println("Server listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, context.ClearHandler(mux)))
}
