package controllers

import (
	"encoding/json"
	"net/http"

	"../auth"
	"../db"
)

func MakeSubTask(w http.ResponseWriter, r *http.Request) {
	var subTask db.SubTask
	err := json.NewDecoder(r.Body).Decode(&subTask)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		panic(err)
	}
	res := subTask.NewSubTask()
	w.Write(res)
}

func EditSubTask(w http.ResponseWriter, r *http.Request) {
	var subTask db.SubTask
	err := json.NewDecoder(r.Body).Decode(&subTask)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		panic(err)
	}
	res := subTask.EditSubTasksByID()
	w.Write(res)
}

func EditStatusSubTask(w http.ResponseWriter, r *http.Request) {
	var subTask db.SubTask
	err := json.NewDecoder(r.Body).Decode(&subTask)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		panic(err)
	}
	res := subTask.EditStatusSubTaskByID()
	w.Write(res)
}

func EditTimeSubTask(w http.ResponseWriter, r *http.Request) {
	var subTask db.SubTask
	err := json.NewDecoder(r.Body).Decode(&subTask)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		panic(err)
	}
	res := subTask.EditTimeSubTaskByID()
	w.Write(res)
}

func GetSubTasks(w http.ResponseWriter, r *http.Request) {
	res, err := db.GetAllSubTasks()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	w.Write(auth.JsonResponseByVar("true", res))
}
