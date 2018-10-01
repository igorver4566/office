package app

import (
	"log"
	"net/http"

	"./routes"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Init() {
	log.Println("Init server...")
	s.port = ":8080"
}

func (s *Server) Start() {
	r := routes.NewRouter()
	log.Println("Starting server")

	http.Handle("/", r.Router)
	http.ListenAndServe(s.port, nil)
}
