package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lottotto/todo-app/api"
	"github.com/lottotto/todo-app/db"
)

func Init() *echo.Echo {
	return getRouter()
}
func getRouter() *echo.Echo {
	e := echo.New()

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler := api.Handler{DB: db.Init()}

	e.GET("/hello", handler.GetHello)
	e.GET("/status", handler.GetHealthCheck)
	e.GET("/task", handler.GetAllTask)
	e.POST("/task", handler.PostTask)
	e.GET("/task/:id", handler.GetTaskById)

	return e
}
