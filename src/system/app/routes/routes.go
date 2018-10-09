package routes

import (
	contr "../controllers"
	"github.com/gorilla/mux"
)

//Router for app
type Router struct {
	Router *mux.Router
}

//NewRouter funcion for create some routes
func NewRouter() (r Router) {
	r.Router = mux.NewRouter()

	r.Router.HandleFunc("/app/register", contr.RegisterHandler).Methods("POST")
	r.Router.HandleFunc("/app/login", contr.LoginHandler)
	return
}
