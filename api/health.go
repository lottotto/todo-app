package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lottotto/todo-app/db"
	"github.com/lottotto/todo-app/model"
)

func GetHealthCheck(c echo.Context) error {
	db := db.Init()

	var health model.Health

	err := db.Ping()

	if err != nil {
		health.Status = http.StatusInternalServerError
		health.Message = "データベース接続エラー"
		return c.JSON(http.StatusInternalServerError, health)
	}
	health.Status = http.StatusOK
	health.Message = "OK"
	return c.JSON(http.StatusOK, health)
}
