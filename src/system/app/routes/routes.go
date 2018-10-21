package routes

import (
	controller "../controllers"
	"github.com/gorilla/mux"
)

//Router for app
type Router struct {
	Router *mux.Router
}

//NewRouter function for create some routes
func NewRouter() (r Router) {
	r.Router = mux.NewRouter()

	r.Router.HandleFunc("/api/register", controller.RegisterHandler).Methods("POST")
	r.Router.HandleFunc("/api/login", controller.LoginHandler).Methods("POST")
	r.Router.HandleFunc("/api/check-token/{token}", controller.CheckToken).Methods("GET")
	r.Router.HandleFunc("/api/tasks/make", controller.MakeTask).Methods("POST")
	r.Router.HandleFunc("/api/tasks/make", controller.GetTaskForm).Methods("GET")
	r.Router.HandleFunc("/api/tasks", controller.GetTasks).Methods("GET")
	return
}
