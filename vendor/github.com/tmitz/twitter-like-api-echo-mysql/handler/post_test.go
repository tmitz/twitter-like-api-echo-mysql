package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/tmitz/twitter-like-api-echo-mysql/model"
)

var (
	postJSON = `{"receive_id": 1, "message": "This is Test."}`
)

type MockHandler struct {
	DB *sqlx.DB
}

func (h *MockHandler) CreatePost(c echo.Context) (err error) {
	u := &model.User{ID: 2}
	p := &model.Post{SendID: u.ID}
	if err = c.Bind(p); err != nil {
		return
	}

	// Validation
	if p.ReceiveID == 0 || p.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid receive_id or send_id fields"}
	}

	if err = h.DB.Get(u, "SELECT * FROM users WHERE id=? LIMIT 1", u.ID); err != nil {
		return echo.ErrNotFound
	}

	tx := h.DB.MustBegin()
	_, err = tx.NamedExec("INSERT INTO posts (receive_id, send_id, message) VALUES (:receive_id, :send_id, :message)", p)
	if err != nil {
		return
	}
	tx.Commit()

	return c.JSON(http.StatusCreated, p)
}

func (h *MockHandler) FetchPost(c echo.Context) (err error) {
	userID := 2
	pageParam := c.QueryParam("page")
	if pageParam == "" {
		pageParam = "0"
	}
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return err
	}

	limitParam := c.QueryParam("limit")
	if limitParam == "" {
		limitParam = "100"
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return err
	}

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit >= 0 {
		limit = 100
	}

	posts := []*model.Post{}
	if err = h.DB.Select(&posts, "SELECT * from posts WHERE receive_id = ? LIMIT ? OFFSET ?", userID, limit, (page-1)*limit); err != nil {
		fmt.Printf("error: %q", err)
		return
	}

	return c.JSON(http.StatusOK, posts)
}

func TestCreatePost(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/posts", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := MockHandler{DB: db}
	err := h.CreatePost(c)
	if err != nil {
		t.Errorf("CreatePost endpoint error: %s", err)
	}

}

func TestFetchPost(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/feed", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := MockHandler{DB: db}
	err := h.FetchPost(c)
	if err != nil {
		t.Errorf("FetchPost endpoint error: %s", err)
	}
}
