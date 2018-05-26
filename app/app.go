package app

import (
	"fmt"
	"net/http"

	"github.com/xiaoheigou/GoFun/config"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}
func (s *Server) Init() {
	fmt.Println("Initialzing server.....")
	s.port = fmt.Sprintf(":%d", config.HttpPort)
}
func (s *Server) Start() {
	fmt.Println("Starting the server listen on ", s.port)
	http.ListenAndServe(s.port, nil)

}
