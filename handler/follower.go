package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tmitz/twitter-like-api-echo-mysql/model/"
)

func (h *Handler) Follow(c echo.Context) (err error) {
	followerID := userIdFromToken(c)
	followedID, _ := strconv.Atoi(c.Param("id"))

	f := &model.Follower{FollowerID: int(followerID), FollowedID: followedID}

	db := h.DB
	tx := db.MustBegin()
	_, err = tx.NamedExec("INSERT INTO followers (follower_id, followed_id) VALUES (:follower_id, :followed_id)", f)
	if err != nil {
		return
	}
	tx.Commit()

	return c.JSON(http.StatusCreated, f)
}
