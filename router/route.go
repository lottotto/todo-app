package router

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lottotto/todo-app/api"
)

func Init(db *sql.DB) *echo.Echo {
	return getRouter(db)
}
func getRouter(db *sql.DB) *echo.Echo {
	e := echo.New()

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler := api.Handler{DB: db}

	e.GET("/hello", handler.GetHello)
	e.GET("/status", handler.GetHealthCheck)
	e.GET("/task", handler.GetAllTask)
	e.POST("/task", handler.PostTask)

	return e
}
