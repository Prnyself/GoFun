package config

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	Cfg       *ini.File
	RunModel  string
	HttpPort  int
	USER_NAME string
	USER_PWD  int
	DATABASE  string
)

func init() {
	var err error
	Cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	loadServer()
	loadDb()
}
func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8080)
}
func loadDb() {
	sec, err := Cfg.GetSection("db")
	if err != nil {
		log.Fatalf("Fail to get section 'db': %v", err)
	}
	USER_NAME = sec.Key("USER_NAME").MustString("root")
	USER_PWD = sec.Key("USER_NAME").MustInt(123456)
	DATABASE = sec.Key("USER_NAME").MustString("func")
}
