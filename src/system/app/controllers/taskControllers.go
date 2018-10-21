package controllers

import (
	"encoding/json"
	"net/http"

	"../auth"
	"../db"
)

func MakeTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		panic(err)
	}
	res := task.NewTask()
	w.Write(res)
}

func GetTaskForm(w http.ResponseWriter, r *http.Request) {
	res := db.GetFormItems()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(auth.JsonResponse(res))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	res, err := db.GetAllTasks()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
	}
	w.Write(auth.JsonResponseByVar("true", res))
}
