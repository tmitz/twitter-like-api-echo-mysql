package model

import (
	"database/sql"
)

type User struct {
	ID       int            `json:"id,omitempty" db:"id"`
	Email    string         `json:"email" db:"email"`
	Password string         `json:"password,omitempty" db:"password"`
	Token    sql.NullString `json:"token,omitempty" db:"token"`
}
