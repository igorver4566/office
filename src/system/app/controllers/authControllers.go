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

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error in request")
		return
	}

	register := user.RegisterUser()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(register)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user db.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error in request")
		return
	}

	login := user.CheckUser()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(login)
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	_, err := auth.IsTokenValid(token)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	w.Write(auth.JsonResponseByVar("true", token))
}
