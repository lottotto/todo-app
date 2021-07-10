package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lottotto/todo-app/db"
	"github.com/lottotto/todo-app/router"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	e := router.Init(db.Init())

	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello Echo", rec.Body.String())
}
func TestGetAllTask(t *testing.T) {

	// あとでmockDBを書く

	e := router.Init(db.Init())
	req := httptest.NewRequest("GET", "/task", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	expected := `[{"id":1,"user_id":1,"type_id":1,"title":"JUnitを学習","detail":"テストの仕方を学習する","deadline":"2020-07-07T15:00:00Z"}]` + "\n"
	actual := rec.Body.String()
	assert.Equal(t, expected, actual)
}
