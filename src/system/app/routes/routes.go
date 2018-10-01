package routes

import (
	contr "../controllers"
	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func NewRouter() (r Router) {

	r.Router = mux.NewRouter()
	r.Router.HandleFunc("/hello", contr.HelloController)
	return
}
