package main

import (
	"github.com/lottotto/todo-app/router"
)

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
