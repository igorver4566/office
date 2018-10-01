package app

import (
	"fmt"
	"log"
	"net/http"

	"./db"
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

	Db := db.InitDb()
	defer Db.Close()
	fmt.Println("Connected")
	insert, err := Db.Query("INSERT INTO Users VALUES ('NULL','Игорь', '12@test.com', '21')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Writed")

	http.Handle("/", r.Router)
	http.ListenAndServe(s.port, nil)
}
