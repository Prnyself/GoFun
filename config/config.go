package config

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	Cfg      *ini.File
	RunModel string
	HttpPort int
)

func init() {
	var err error
	Cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	loadServer()
}
func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8080)
}
