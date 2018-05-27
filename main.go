package main

import (
	"fmt"
	"log"

	"github.com/xiaoheigou/GoFun/app"
)

func main() {
	s := app.NewServer()
	s.Init()

	result, err := s.DB.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("table:%v", result)

	s.Start()
}
