package controllers

import (
	"net/http"

	"../auth"
	"../db"
	"github.com/gorilla/mux"
)

func GetAllWorkers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	start := vars["dt_start"]
	end := vars["dt_end"]
	res, err := db.GetAllWorkers(start, end)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	w.Write(auth.JsonResponseByVar("true", res))
}
