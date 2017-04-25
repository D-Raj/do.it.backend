package models

import "database/sql"

// Action - an action performed by a user on a particular date
type Action struct {
	UserID   int    `json:"user_id"`   // actions.user_id
	GoalID   int    `json:"goal_id"`   // actions.goal_id
	Date     int64  `json:"date"`      // actions.date
	Weight   int    `json:"weight"`    // actions.weight
	GoalName string `json:"goal_name"` // goals.name
	Achieved int    `json:"achieved"`  // actions.achieved
}

// GetActions - get Actions (actions/goals in db) for a given user
func GetActions(userID int, date int64) ([]*Action, error) {
	query := `
                SELECT
                        actions.user_id,
                        actions.goal_id,
                        actions.day as date,
                        goals.name,
                        actions.weight,
                        actions.achieved
                FROM actions
                INNER JOIN goals
                ON actions.goal_id = goals.id
                WHERE actions.user_id = ?
        `

	var err error
	var rows *sql.Rows

	if date == 0 {
		rows, err = db.Query(query, userID)
	} else {
		query += `AND actions.day = ?`
		rows, err = db.Query(query, userID, date)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	actions := make([]*Action, 0)
	for rows.Next() {
		action := new(Action)
		err := rows.Scan(
			&action.UserID,
			&action.GoalID,
			&action.Date,
			&action.GoalName,
			&action.Weight,
			&action.Achieved)
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
func CreateAction(action Action) (int, error) {
	query := `
                INSERT INTO actions
                        (user_id, goal_id, day, weight, achieved)
                VALUES (?, ?, ?, ?, ?)
                ON DUPLICATE KEY UPDATE
                        user_id = ?,
                        goal_id = ?,
                        day = ?,
                        weight = ?,
                        achieved = ?;
        `
	_, err := db.MustExec(
		query,
		action.UserID,
		action.GoalID,
		action.Date,
		action.Weight,
		action.Achieved,
		action.UserID,
		action.GoalID,
		action.Date,
		action.Weight,
		action.Achieved).LastInsertId()
	if err != nil {
		return 0, err
	}
	return 1, nil
}
