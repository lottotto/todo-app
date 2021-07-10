package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lottotto/todo-app/api"
)

func Init() *echo.Echo {
	return getRouter()
}

func getRouter() *echo.Echo {
	e := echo.New()

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", api.GetHello)
	e.GET("/task", api.GetAllTask)
	e.GET("/status", api.GetHealthCheck)

	return e
}
