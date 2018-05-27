package app

import (
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	"github.com/xiaoheigou/GoFun/config"
	"github.com/xiaoheigou/GoFun/pkg/db"

	"github.com/xiaoheigou/GoFun/routers"
)

type Server struct {
	Port string
	DB   *xorm.Engine
}

func NewServer() Server {
	return Server{}
}
func (s *Server) Init() {
	var err error
	fmt.Println("Initialzing server.....")
	s.Port = fmt.Sprintf(":%d", config.HttpPort)
	s.DB, err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
func (s *Server) Start() {
	fmt.Println("Starting the server listen on ", s.Port)
	router := routers.InitRouter()
	router.Run(s.Port)
}
