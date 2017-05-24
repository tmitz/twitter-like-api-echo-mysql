package handler

import (
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	var err error
	db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:13306)/test_twitter")
	if err != nil {
		log.Fatalf("Failed to db connection: %s", err)
	}
	defer db.Close()

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

	db.MustExec(
		"INSERT INTO users (id, email, password) VALUES (1, ?, ?), (2, ?, ?)",
		"abc@example.com",
		"foofoo",
		"xyz@example.com",
		"barbar",
	)

	os.Exit(m.Run())
}
