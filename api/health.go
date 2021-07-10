package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lottotto/todo-app/model"
)

func (hander Handler) GetHealthCheck(c echo.Context) error {

	var health model.Health

	err := hander.DB.Ping()

	if err != nil {
		health.Status = http.StatusInternalServerError
		health.Message = "データベース接続エラー"
		return c.JSON(http.StatusInternalServerError, health)
	}
	health.Status = http.StatusOK
	health.Message = "OK"
	return c.JSON(http.StatusOK, health)
}
