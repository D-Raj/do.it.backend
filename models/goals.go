package models

import "database/sql"

// Goal - a goal as seen in db
type Goal struct {
	ID     int    `json:"id"`     // goals.id
	Name   string `json:"name"`   // goals.value
	Weight int    `json:"weight"` // actions.weight
}

// GetActiveGoals - get currently-active goals for a given user
func GetActiveGoals(userID int) ([]*Goal, error) {
	query := `
                SELECT goals.id, goals.name, active_goals.weight
                FROM active_goals
                INNER JOIN goals
                ON goals.id = active_goals.goal_id
                WHERE active_goals.user_id = ?;
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
			&goal.Name,
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
	goalID, err := getGoalID(value)
	if err != nil {
		return 0, err
	}

	query := `
                INSERT INTO active_goals (user_id, goal_id, weight) VALUES (?, ?, ?)
                ON DUPLICATE KEY UPDATE weight = ?
        `
	result, err := db.Exec(query, userID, goalID, weight, weight)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(insertID), nil
}

func getGoalID(value string) (int, error) {
	query := `SELECT id FROM goals WHERE value = ?`
	var goalID int
	err := db.QueryRowx(query, value).Scan(&goalID)
	if err == nil {
		return int(goalID), nil
	} else if err != sql.ErrNoRows {
		return 0, err
	}

	query = `INSERT INTO goals (value) VALUES (?)`
	result, err := db.Exec(query, value)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}
