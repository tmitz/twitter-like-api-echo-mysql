package handler

import (
	"database/sql"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/tmitz/twitter-like-api-echo-mysql/model"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind
	u := &model.User{}
	if err = c.Bind(u); err != nil {
		return
	}

	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	tx := h.DB.MustBegin()
	_, err = tx.NamedExec("INSERT INTO users (email, password) VALUES (:email, :password)", u)
	if err != nil {
		return
	}
	tx.Commit()

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}
	// Select
	if err = h.DB.Get(u, "SELECT id, email, password, token FROM users WHERE email=? AND password=? LIMIT 1", u.Email, u.Password); err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
	}

	//-----
	// JWT
	//-----

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenValue, err := token.SignedString([]byte(Key))
	if err != nil {
		return
	}
	u.Token = sql.NullString{Valid: true, String: tokenValue}
	u.Password = ""
	return c.JSON(http.StatusOK, u)
}

func userIdFromToken(c echo.Context) float64 {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(float64)
}
