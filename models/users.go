package models

import "database/sql"

// User - single user record in db
type User struct {
	externalID   string
	userSourceID int
	name         string
	email        string
}

// AddUser - insert user if not exist
func AddUser(sourceName string, externalID string, name string, email string) (int, error) {
	userSourceID, err := getUserSourceID(sourceName)
	if err != nil {
		return 0, err
	}

	user := User{
		externalID:   externalID,
		userSourceID: userSourceID,
		name:         name,
		email:        email,
	}

	insertID, err := db.MustExec(
		"INSERT INTO users (external_id, user_source_id, name, email) VALUES (?, ?, ?, ?)",
		user.externalID,
		user.userSourceID,
		user.name,
		user.email).LastInsertId()

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
		// insert and get new id if there was no source result
		id, err = db.MustExec("INSERT INTO user_sources (name) VALUES (?)", sourceName).LastInsertId()
	} else if err != nil {
		return 0, err
	}

	return int(id), nil
}
