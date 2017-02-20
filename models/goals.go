package models

import (
	"database/sql"
	"fmt"
)

// User - single user as represented by client (not a db record)
type User struct {
	ID    int    // internal id for a user
	Name  string // name of user
	Email string // email of user
}

// GetUserID - get user internal id, insert if not found and
func GetUserID(sourceName string, externalID string, name string, email string) (int, error) {
	userSourceID, err := getUserSourceID(sourceName)
	if err != nil {
		return 0, err
	}

	var userID int
	err = db.QueryRowx("SELECT id FROM users WHERE external_id = ? and user_source_id = ?", externalID, userSourceID).Scan(&userID)
	if err == nil {
		return userID, nil
	} else if err != sql.ErrNoRows {
		return 0, err
	}

	userRecord := struct {
		userSourceID int
		externalID   string
		name         string
		email        string
	}{
		userSourceID, externalID, name, email,
	}

	result, err := db.Exec(
		"INSERT INTO users (external_id, user_source_id, name, email) VALUES (?, ?, ?, ?)",
		userRecord.externalID,
		userRecord.userSourceID,
		userRecord.name,
		userRecord.email,
	)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(insertID), nil
}

func getUserSourceID(sourceName string) (int, error) {
	// try to get source id
	var id int64
	err := db.QueryRowx("SELECT id FROM user_sources WHERE name = ?", sourceName).Scan(&id)
	if err == sql.ErrNoRows {
		fmt.Println("no id?")
		fmt.Println(err)
		// insert and get new id if there was no source result
		id, err = db.MustExec("INSERT INTO user_sources (name) VALUES (?)", sourceName).LastInsertId()
	} else if err != nil {
		fmt.Println("wtf")
		fmt.Println(err)
		return 0, err
	}

	return int(id), nil
}
