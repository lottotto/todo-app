package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lottotto/todo-app/model"
)

// ここなんか一つ上のレイヤーっぽいが、、、
type Handler struct {
	DB *sql.DB
}

func (hander Handler) GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Echo")
}

func (hander Handler) GetAllTask(c echo.Context) error {
	// elastic apm用に修正
	rows, err := hander.DB.QueryContext(c.Request().Context(), "SELECT id, user_id, type_id, title, detail, deadline, done from task")

	if err != nil {
		panic(err)
	}

	var taskResult []model.Task
	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.Id, &task.UserId, &task.TypeId, &task.Title, &task.Detail, &task.Deadline, &task.Done)

		if err != nil {
			panic(err)
		}

		taskResult = append(taskResult, task)
	}
	return c.JSON(http.StatusOK, taskResult)
}

func (hander Handler) PostTask(c echo.Context) error {
	var task model.Task
	if err := c.Bind(&task); err != nil {
		panic(err)
	}
	// rowsの戻りがない場合は db.Exec() を使うこと。postgresは戻りがあるので、queryでOK
	// 戻りのIDを取得するには、QueryRowはないとダメ、queryだとうごきません
	err := hander.DB.QueryRowContext(c.Request().Context(), "INSERT INTO task(user_id, type_id, title, detail, deadline, done) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;",
		task.UserId, task.TypeId, task.Title, task.Detail, task.Deadline).Scan(&task.Id)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusAccepted, task)
}
func (handler Handler) GetTaskById(c echo.Context) error {
	var task model.Task
	query := `select id, user_id, type_id, title, detail, deadline, done from task where id=$1;`
	rows, err := handler.DB.QueryContext(c.Request().Context(), query, c.Param("id"))

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&task.Id, &task.UserId, &task.TypeId, &task.Title, &task.Detail, &task.Deadline, &task.Done)
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, task)
}
