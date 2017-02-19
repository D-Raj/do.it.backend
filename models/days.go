package models

import "database/sql"

// Day - score for a user on a given day
type Day struct {
	Date  int64   `json:"date"`  // actions.day (grouped by)
	Score float32 `json:"score"` // calculated score (percentage done)
}

// GetAllDays - get scores for each day for a given user
func GetAllDays(userID int) ([]*Day, error) {
	query := `
                        SELECT a.day as date, (d_weight / a_weight) AS score
                        FROM (
                                SELECT day,
                                SUM(weight) AS d_weight
                                FROM actions
                                WHERE achieved = true
                                AND user_id = ?
                                GROUP BY day
                        ) AS d
                        INNER JOIN (
                                SELECT day,
                                SUM(weight) AS a_weight
                                FROM actions
                                WHERE user_id = ?
                                GROUP BY day
                        ) AS a
                        ON d.day = a.day
        `
	rows, err := db.Query(query, userID, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	days := make([]*Day, 0)
	for rows.Next() {
		day := new(Day)
		err := rows.Scan(
			&day.Date,
			&day.Score)
		if err != nil {
			return nil, err
		}
		days = append(days, day)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return days, nil
}
