package models

import "database/sql"

// Goal - a goal as seen in db
type Goal struct {
	ID     int    `json:"id"`     // goals.id
	Value  string `json:"value"`  // goals.value
	Weight int    `json:"weight"` // goals.weight
}

// GetAllGoals - get goals (as seen in db) for a given user
func GetAllGoals(userID int) ([]*Goal, error) {
	query := `
                SELECT goals.*
                FROM goals
                INNER JOIN users_goals
                        ON users_goals.goal_id = goals.id
                INNER JOIN users
                        ON users.id = users_goals.user_id
                WHERE users.id = ?
        `
	rows, err := db.Queryx(query, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	goals := make([]*Goal, 0)
	for rows.Next() {
		goal := new(Goal)
		err := rows.Scan(
			&goal.ID,
			&goal.Value,
			&goal.Weight)
		if err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return goals, nil
}

// CreateGoal - create an goal for a logged in user
func CreateGoal(userID int, value string, weight int) (int, error) {
	goalID, err := getGoalID(value, weight)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO users_goals (user_id, goal_id) VALUES (?, ?)`
	result, err := db.Exec(query, userID, goalID)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(insertID), nil
}

func getGoalID(value string, weight int) (int, error) {
	query := `SELECT id FROM goals WHERE value = ? and weight = ?`
	var goalID int
	err := db.QueryRowx(query, value, weight).Scan(&goalID)
	if err == nil {
		return int(goalID), nil
	} else if err != sql.ErrNoRows {
		return 0, err
	}

	query = `INSERT INTO goals (value, weight) VALUES (?, ?)`
	result, err := db.Exec(query, value, weight)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}
