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
	apiRouter.GET("/api/me/actions", AllActionsHandler)
	apiRouter.POST("/api/me/actions", NewActionHandler)

	/* days read (aggregate activity for the past year) */
	apiRouter.GET("/api/me", DaysHandler)

	loggerHandler := middleware.NewLoggerHandler(apiRouter)
	setHeadersHandler := middleware.NewSetHeadersHandler(loggerHandler)
	apiHandler := middleware.NewAuthHandler(setHeadersHandler, sessionName, store)

	authHandler := middleware.NewAuthHandler(authRouter, sessionName, store)

	mux := http.NewServeMux()
	mux.Handle("/auth/", authHandler)
	mux.Handle("/api/", apiHandler)
	mux.HandleFunc("/", StaticHandler)

	fmt.Println("Server listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, context.ClearHandler(mux)))
}
