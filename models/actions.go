package models

import (
	"database/sql"
	"fmt"
)

// Action - an action performed by a user on a particular day
type Action struct {
	UserID     int    `json:"user_id"`     // actions.user_id
	Day        string `json:"day"`         // actions.day
	GoalID     int    `json:"goal_id"`     // goals.id
	GoalValue  string `json:"goal_value"`  // goals.value
	GoalWeight int    `json:"goal_weight"` // goals.weight
}

// GetAllActions - get Actions (actions/goals in db) for a given user
func GetAllActions(userID int) ([]*Action, error) {
	query := `
                SELECT
                        actions.user_id, actions.day, goals.id, goals.value, goals.weight
                FROM actions INNER JOIN goals
                        ON actions.goal_id = goals.id
                WHERE actions.user_id = ?
        `
	rows, err := db.Query(query, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	actions := make([]*Action, 0)
	for rows.Next() {
		action := new(Action)
		err := rows.Scan(
			&action.UserID,
			&action.Day,
			&action.GoalID,
			&action.GoalValue,
			&action.GoalWeight)
		if err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return actions, nil
}

// CreateAction - create an action for a logged in user
func CreateAction(userID int, goalID int, day int) (int, error) {
	fmt.Printf("userID: %d, goalID: %d, day: %d\n\n\n", userID, goalID, day)
	query := `
                INSERT INTO actions
                        (user_id, goal_id, day)
                VALUES (?, ?, ?)
        `
	id, err := db.MustExec(query, userID, goalID, day).LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
