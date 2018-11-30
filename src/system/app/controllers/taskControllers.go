package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../auth"
	"../db"
	"github.com/gorilla/mux"
)

func MakeTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
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
		return
	}
	w.Write(auth.JsonResponseByVar("true", res))
}

func GetOneTask(w http.ResponseWriter, r *http.Request) {
	var res struct {
		Task db.TaskReturn      `json:"task"`
		Sub  []db.SubTaskReturn `json:"items"`
	}

	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.ParseInt(id, 10, 64)
	task, err := db.GetTaskById(idInt)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	subs, err := db.GetSubTasksByTask(strconv.FormatInt(int64(task.ID), 10))
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	res.Task = task
	res.Sub = subs
	w.Write(auth.JsonResponseByVar("true", res))
}
