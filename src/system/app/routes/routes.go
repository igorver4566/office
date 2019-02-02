package routes

import (
	"fmt"
	"log"
	"net/http"

	auth "../auth"
	controller "../controllers"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
)

//Router for app
type Router struct {
	Router *mux.Router
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("ts")
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return auth.VerifyKey, nil
			})

		if err == nil {

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Token is not valid")
				log.Println("Token is not valid")
			}
		} else {
			if r.URL.Path == "/api/login" || r.URL.Path == "/api/register" {
				log.Println("get")
				next.ServeHTTP(w, r)
			} else {

				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Unauthorised access to this resource")

			}

		}
	})
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
	r.Router.HandleFunc("/api/subtasks/make", controller.MakeSubTask).Methods("POST")
	r.Router.HandleFunc("/api/subtasks/edit", controller.EditSubTask).Methods("POST")
	r.Router.HandleFunc("/api/subtasks/editStatus", controller.EditStatusSubTask).Methods("POST")
	r.Router.HandleFunc("/api/subtasks/editTime", controller.EditTimeSubTask).Methods("POST")
	r.Router.HandleFunc("/api/subtasks", controller.GetSubTasks).Methods("GET")
	r.Router.HandleFunc("/api/task/{id}", controller.GetOneTask).Methods("GET")
	r.Router.HandleFunc("/api/chat/{name}", controller.GetChatHistory).Methods("GET")
	r.Router.HandleFunc("/api/chat/{name}", controller.PostMessage).Methods("POST")
	r.Router.HandleFunc("/api/users/workers/{dt_start}/{dt_end}", controller.GetAllWorkers).Methods("GET")
	r.Router.Use(LoggingMiddleware)
	return
}
