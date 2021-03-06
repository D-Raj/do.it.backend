package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"bh/do.it/models"

	"github.com/jinzhu/now"
	"github.com/julienschmidt/httprouter"
)

var secondsInDay = 86400
var daysInYear = 365

// DaysHandler - get scores for each day for the past year with a logged in user
func DaysHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, err := getUserID(r)
	if err != nil {
		HandleError(err, w)
		return
	}

	days, err := models.GetAllDays(userID)
	if err != nil {
		HandleError(err, w)
		return
	}

	// generate year of empty days
	today := now.BeginningOfDay().Unix()
	year := make([]models.Day, 0)
	indexLookup := make(map[int64]int)
	for i := 0; i < (daysInYear - 1); i++ {
		date := today - int64((i * secondsInDay))
		day := models.Day{
			Date:  date,
			Score: 0,
		}
		year = append(year, day)
		indexLookup[date] = i
	}

	// populate the year with results from user days results
	for _, day := range days {
		i := indexLookup[day.Date]
		year[i].Score = day.Score
	}

	response, err := json.Marshal(year)
	if err != nil {
		HandleError(err, w)
		return
	}

	w.Write(response)
}

// ActionsHandler - get all actions associated with logged in user on a particular day
func ActionsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, err := getUserID(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	date, err := getDate(r, 0)
	if err != nil {
		HandleError(err, w)
		return
	}

	actions, err := models.GetActions(userID, date)
	if err != nil {
		HandleError(err, w)
		return
	}

	response, err := json.Marshal(actions)
	if err != nil {
		HandleError(err, w)
		return
	}

	w.Write(response)
}

// NewActionHandler - create an action given a goal, user, and optional day (unix start of day)
func NewActionHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, err := getUserID(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	goalID, err := getGoalID(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	date, err := getDate(r, -1)
	if err != nil {
		HandleError(err, w)
		return
	}
	weight, err := getWeight(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	achieved, err := getAchieved(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	name := ""

	action := models.Action{
		UserID:   userID,
		GoalID:   goalID,
		Date:     date,
		Weight:   weight,
		GoalName: name,
		Achieved: achieved,
	}

	_, err = models.CreateAction(action)
	if err != nil {
		HandleError(err, w)
		return
	}

	w.Write([]byte("{\"success\":\"ok\"}"))
}

// ActiveGoalsHandler - get all goals associated with logged in user
func ActiveGoalsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, err := getUserID(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	goals, err := models.GetActiveGoals(userID)
	if err != nil {
		HandleError(err, w)
		return
	}

	response, err := json.Marshal(goals)
	if err != nil {
		HandleError(err, w)
		return
	}

	w.Write(response)
}

// NewGoalHandler - create a goal given a goal, user, and optional day (unix start of day)
func NewGoalHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	name, err := getName(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	weight, err := getWeight(r)
	if err != nil {
		HandleError(err, w)
		return
	}

	_, err = models.CreateGoal(userID, name, weight)
	if err != nil {
		HandleError(err, w)
		return
	}

	w.Write([]byte("{\"success\":\"ok\"}"))
}

func getUserID(r *http.Request) (int, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return 0, err
	}

	val := session.Values["user"]
	user, ok := val.(*models.User)
	if !ok {
		return 0, err
	}

	return user.ID, nil
}

func getGoalID(r *http.Request) (int, error) {
	goalString := r.FormValue("goal_id")
	goalID, err := strconv.ParseInt(goalString, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(goalID), nil
}

func getDate(r *http.Request, defaultValue int64) (int64, error) {
	dateString := r.FormValue("date")
	if dateString == "" {
		if defaultValue == -1 {
			date := now.BeginningOfDay().Unix()
			return date, nil
		}
		date := defaultValue
		return date, nil
	}
	date, err := strconv.ParseInt(dateString, 10, 32)
	if err != nil {
		return 0, nil
	}
	return date, nil
}

func getName(r *http.Request) (string, error) {
	value := r.FormValue("value")
	if value == "" {
		return "", errors.New("Invalid goals value")
	}
	return value, nil
}

func getWeight(r *http.Request) (int, error) {
	weightString := r.FormValue("weight")
	weight, err := strconv.ParseInt(weightString, 10, 32)
	if err != nil {
		return 0, err
	}

	return int(weight), nil
}

func getAchieved(r *http.Request) (int, error) {
	achievedString := r.FormValue("achieved")
	if achievedString == "true" {
		return 1, nil
	}
	if achievedString == "false" {
		return 0, nil
	}
	return 0, errors.New("Invalid achieved value")
}
