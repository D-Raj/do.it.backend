package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"bh/do.it/models"

	"strconv"

	"github.com/jinzhu/now"
	"github.com/julienschmidt/httprouter"
)

// AllActionsHandler - get all actions associated with logged in user
func AllActionsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, err := getUserID(r)
	if err != nil {
		HandleError(err, w)
		return
	}

	actions, err := models.GetAllActions(userID)
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
	day, err := getDay(r)
	if err != nil {
		HandleError(err, w)
		return
	}

	_, err = models.CreateAction(userID, goalID, day)
	if err != nil {
		HandleError(err, w)
		return
	}

	w.Write([]byte("{\"success\":\"ok\"}"))
}

// AllGoalsHandler - get all goals associated with logged in user
func AllGoalsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, err := getUserID(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	goals, err := models.GetAllGoals(userID)
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
	value, err := getValue(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	weight, err := getWeight(r)
	if err != nil {
		HandleError(err, w)
		return
	}

	_, err = models.CreateGoal(userID, value, weight)
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

func getDay(r *http.Request) (int, error) {
	dayString := r.FormValue("day")
	if dayString == "" {
		day := now.BeginningOfDay().Unix()
		return int(day), nil
	}
	day, err := strconv.ParseInt(dayString, 10, 32)
	if err != nil {
		return 0, nil
	}
	return int(day), nil
}

func getValue(r *http.Request) (string, error) {
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
