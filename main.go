package main

import (
	"fmt"
	"log"
	"net/http"

	"bh/do.it/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

var db *sqlx.DB

func main() {
	models.InitDB("do_it_dev:do_it_dev@/do_it_dev")

	router := httprouter.New()

	/* authentication */
	router.GET("/", DebugIndex) // TEMP for dev/debug
	router.GET("/GoogleLogin", HandleGoogleLogin)
	router.GET("/GoogleCallback", HandleGoogleCallback)

	/* actions/goals */
	// router.GET("/me", allActionGoals)
	// router.POST("/me", newAction)

	/* goals */
	// router.GET("me/goals", allGoals)
	// router.POST("me/goals", newGoal)

	fmt.Println("Server listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

// DebugIndex - comment for lint shut up
// func DebugIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	models.AddUser("123456789", "NEW_SOURCE", "Leaf Suarez", "leaf@suarez.net")
// }
