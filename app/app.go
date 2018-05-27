package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/xiaoheigou/GoFun/config"
	"github.com/xiaoheigou/GoFun/db"
)

type Server struct {
	port string
	DB   *xorm.Engine
}

func NewServer() Server {
	return Server{}
}
func (s *Server) Init() {
	var err error
	fmt.Println("Initialzing server.....")
	s.port = fmt.Sprintf(":%d", config.HttpPort)
	s.DB, err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
func (s *Server) Start() {
	fmt.Println("Starting the server listen on ", s.port)
	http.ListenAndServe(s.port, nil)

}
