package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../auth"
	"../db"
	"github.com/gorilla/mux"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := db.UserCredential()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error in request")
		return
	}

	token, id, err := user.RegisterUser()
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	res, err := db.GetUserById(int64(id), token)
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}

	w.Write(auth.JsonResponseByVar("true", res))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user db.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(auth.JsonResponseByVar("false", "Ошибка в запросе"))
		return
	}
	token, id, err := user.CheckUser()
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	res, err := db.GetUserById(int64(id), token)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(auth.JsonResponseByVar("true", res))
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	id, err := auth.IsTokenValid(token)
	if err != nil || id == 0 {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	res, err := db.GetUserById(id, token)
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	w.Write(auth.JsonResponseByVar("true", res))
}
