package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tmitz/twitter-like-api-echo-mysql/model"
)

func (h *Handler) CreatePost(c echo.Context) (err error) {
	u := &model.User{ID: int(userIdFromToken(c))}
	p := &model.Post{SendID: u.ID}
	if err = c.Bind(p); err != nil {
		return
	}

	// Validation
	if p.ReceiveID == 0 || p.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid receive_id or send_id fields"}
	}

	db := h.DB
	if err = db.Get(u, "SELECT * FROM users WHERE id=? LIMIT 1", u.ID); err != nil {
		return echo.ErrNotFound
	}

	tx := db.MustBegin()
	_, err = tx.NamedExec("INSERT INTO posts (receive_id, send_id, message) VALUES (:receive_id, :send_id, :message)", p)
	if err != nil {
		return
	}
	tx.Commit()

	return c.JSON(http.StatusCreated, p)
}

func (h *Handler) FetchPost(c echo.Context) (err error) {
	userID := int(userIdFromToken(c))
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return err
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
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
	db := h.DB
	if err = db.Select(posts, "SELECT * from posts WHERE receive_id = ? LIMIT ? OFFSET ?", userID, limit, (page-1)*limit); err != nil {
		return
	}

	return c.JSON(http.StatusOK, posts)
}
