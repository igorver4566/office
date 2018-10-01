package app

import (
	"log"
	"net/http"
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
	log.Println("Starting server")
	http.ListenAndServe(s.port, nil)
}
