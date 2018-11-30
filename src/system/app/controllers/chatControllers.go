package controllers

import (
	"encoding/json"
	"net/http"

	"../auth"
	"../slack"
	"github.com/gorilla/mux"
)

func GetChatHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	messages, err := slack.ChatHistory(name)
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	w.Write(auth.JsonResponseByVar("true", messages))

}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var res struct {
		Msg  string `json:"msg"`
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&res)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		panic(err)
	}
	_, err = slack.PostMessage(name, res.Msg, res.Name)
	if err != nil {
		w.Write(auth.JsonResponseByVar("false", err.Error()))
		return
	}
	w.Write(auth.JsonResponseByVar("true", "Сообщение отправлено"))
}
