package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// InitDB - open db connection and ping to make sure it's ready
func InitDB(dataSourceName string) {
	var err error
	db, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
