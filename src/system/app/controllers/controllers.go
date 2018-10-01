package controllers

import (
	"fmt"
	"net/http"
)

func HelloController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
