package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/lottotto/todo-app/api"
	"github.com/lottotto/todo-app/model"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestHelloHandler(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	h := api.Handler{}
	if assert.NoError(t, h.GetHello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello Echo", rec.Body.String())
	}
}
func TestGetAllTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest("GET", "/task", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sql := `SELECT id, user_id, type_id, title, detail, deadline from task`
	columns := []string{"id", "user_id", "type_id", "title", "detail", "deadline"}
	mock.ExpectQuery(sql).WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, 1, "JUnitを学習", "テストの仕方を学習する", time.Date(2020, time.July, 7, 15, 0, 0, 0, time.UTC)))

	task := model.Task{Id: 1, UserId: 1, TypeId: 1, Title: "JUnitを学習", Detail: "テストの仕方を学習する", Deadline: time.Date(2020, time.July, 7, 15, 0, 0, 0, time.UTC)}
	var tasks []model.Task
	tasks = append(tasks, task)
	expected, err := json.Marshal(tasks)
	if err != nil {
		t.Fatal("json marsgal error")
	}

	h := api.Handler{DB: db}
	if assert.NoError(t, h.GetAllTask(c)) {
		actual := rec.Body.String()
		assert.Equal(t, http.StatusOK, rec.Code)
		// httpの仕様により改行が入る
		assert.Equal(t, string(expected)+"\n", actual)
	}
}

func TestPostTask(t *testing.T) {

	// request までのを作成する
	e := echo.New()
	requestParamTask := model.Task{
		UserId:   1,
		TypeId:   1,
		Title:    "Go言語を学習",
		Detail:   "Go言語のWEBアプリケーションフレームワークを学習する",
		Deadline: time.Date(2021, time.August, 11, 15, 0, 0, 0, time.UTC),
	}
	requestParam, err := json.Marshal(requestParamTask)
	if err != nil {
		t.Fatal("json marsgal error")
	}
	req := httptest.NewRequest("POST", "/task", bytes.NewReader(requestParam))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//dbmock を設定
	db, mock, err := sqlmock.New()
	sql := `INSERT INTO task(.+) VALUES (.+) RETURNING id;`
	mock.ExpectQuery(sql).WithArgs(
		requestParamTask.UserId,
		requestParamTask.TypeId,
		requestParamTask.Title,
		requestParamTask.Detail,
		requestParamTask.Deadline,
	).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	h := api.Handler{DB: db}

	if assert.NoError(t, h.PostTask(c)) {
		assert.Equal(t, http.StatusAccepted, rec.Code)
	}
}

func TestGetTaskById(t *testing.T) {
	// setting echo
	e := echo.New()
	req := httptest.NewRequest("GET", "/task/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// sql := `select id, user_id, type_id, title, detail, deadline from task where id=(.+);`
	sql := `select (.+) from task where id=(.+);`

	columns := []string{"id", "user_id", "type_id", "title", "detail", "deadline"}

	mock.ExpectQuery(sql).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1,
				1,
				1,
				"JUnitを学習",
				"テストの仕方を学習する",
				time.Date(2020, time.July, 7, 15, 0, 0, 0, time.UTC),
			),
		)
	h := api.Handler{DB: db}
	task := model.Task{
		Id:       1,
		UserId:   1,
		TypeId:   1,
		Title:    "JUnitを学習",
		Detail:   "テストの仕方を学習する",
		Deadline: time.Date(2020, time.July, 7, 15, 0, 0, 0, time.UTC),
	}
	expected, err := json.Marshal(task)
	if err != nil {
		t.Fatal("json marshal error")
	}

	if assert.NoError(t, h.GetTaskById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expected)+"\n", rec.Body.String())

	}

}
