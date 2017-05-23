package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/tmitz/twitter-like-api-echo-mysql/handler"
)

var db *sqlx.DB

var (
	userJSON = "{\"email\":\"jon@example.com\", \"password\":\"foobarbaz\"}"
)

func TestMain(m *testing.M) {
	mockdb, err := sqlx.Open("mysql", "root@tcp(127.0.0.1:13306)/test_twitter")
	if err != nil {
		log.Fatalf("Failed to db connection: %s", err)
	}
	db = mockdb
	defer mockdb.Close()

	tables, err := db.Queryx("SHOW TABLES")
	if err != nil {
		log.Fatalf("SHOW TABLES error: %s", err)
	}
	defer tables.Close()

	db.MustExec("SET FOREIGN_KEY_CHECKS = 0")
	for tables.Next() {
		var table string
		err = tables.Scan(&table)
		if err != nil {
			log.Fatalf("SHOW TABLES scan error: %s", err)
		}
		db.MustExec("TRUNCATE TABLE " + table)
	}
	db.MustExec("SET FOREIGN_KEY_CHECKS = 1")

	os.Exit(m.Run())
}

func TestSignup(t *testing.T) {
	fmt.Printf("db status2: %q\n", db)
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
		t.Errorf("invalid http status code: %s, got: %s", http.StatusCreated, rec.Code)
	}
}
