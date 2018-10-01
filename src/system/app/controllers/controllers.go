package controllers

import (
	"fmt"
	"net/http"

	"../db"
)

func HelloController(w http.ResponseWriter, r *http.Request) {
	s := db.Insert("test", "test", "2")
	fmt.Fprint(w, s)
}
