package app

import (
	"fmt"
	"net/http"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}
func (s *Server) Init() {
	fmt.Println("Initialzing server.....")
	s.port = ":8080"
}
func (s *Server) Start() {
	fmt.Println("Starting the server listen on ", s.port)
	http.ListenAndServe(s.port, nil)

}
