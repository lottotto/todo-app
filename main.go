package main

import (
	"github.com/lottotto/todo-app/db"
	"github.com/lottotto/todo-app/router"
)

func main() {
	e := router.Init(db.Init())
	e.Logger.Fatal(e.Start(":1323"))
}
