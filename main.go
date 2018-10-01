package main

import (
	"./src/system/app"
)

func main() {
	s := app.NewServer()

	s.Init()
	s.Start()
}
