package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/tmitz/twitter-like-api-echo-mysql/handler"
)

var (
	userJSON = "{\"email\":\"jon@example.com\", \"password\":\"foobarbaz\"}"
)

func TestSignup(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler.Handler{DB: db}
	err := h.Signup(c)
	if err != nil {
		t.Errorf("Cause signup endpoint error: %s", err)
	}
	if http.StatusCreated != rec.Code {
		t.Errorf("invalid http status code: %d, got: %d", http.StatusCreated, rec.Code)
	}
}

func TestLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/login", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler.Handler{DB: db}
	err := h.Login(c)
	if err != nil {
		t.Errorf("Login endpoint error: %s", err)
	}
	if http.StatusOK != rec.Code {
		t.Errorf("Invalid http status code: %d, got: %d", rec.Code, http.StatusOK)
	}
}
