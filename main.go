package main

import (
	"github.com/xiaoheigou/GoFun/app"
)

func main() {
	s := app.NewServer()
	s.Init()

	s.Start()

}
