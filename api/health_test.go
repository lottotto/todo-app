package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/lottotto/todo-app/api"
	"github.com/lottotto/todo-app/db"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest("GET", "/status", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	dbmock, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := api.Handler{dbmock}
	if assert.NoError(t, h.GetHealthCheck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestHealthCheckNoConnect(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest("GET", "/status", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	h := api.Handler{DB: db.Init()}
	err := h.GetHealthCheck(c)
	if err != nil {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
