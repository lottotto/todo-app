package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lottotto/todo-app/db"
	"github.com/lottotto/todo-app/model"
)

func GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Echo")
}

func GetAllTask(c echo.Context) error {
	db := db.Init()
	rows, err := db.Query("SELECT id, user_id, type_id, title, detail, deadline from task")

	if err != nil {
		panic(err)
	}

	var taskResult []model.Task
	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.Id, &task.UserId, &task.TypeId, &task.Title, &task.Detail, &task.Deadline)

		if err != nil {
			panic(err)
		}

		taskResult = append(taskResult, task)
	}
	return c.JSON(http.StatusOK, taskResult)
}
