package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"../db"
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

	fmt.Println(user.Login, user.Password)

	if strings.ToLower(user.Login) != "alexcons" {
		if user.Password != "kappa123" {
			fmt.Println("Error logging in")
			return
		}
	}
}
