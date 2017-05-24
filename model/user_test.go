package model

import (
	"database/sql"
	"encoding/json"
	"testing"
)

func TestUserWithoutToken(t *testing.T) {
	want := "{\"email\":\"abc@example.com\",\"password\":\"awwwww\",\"token\":{\"String\":\"\",\"Valid\":false}}"
	u := User{Email: "abc@example.com", Password: "awwwww"}

	json, err := json.Marshal(u)
	if err != nil {
		t.Errorf("json marshal error: %s", err)
	}
	if string(json) != want {
		t.Errorf("Marshaled json expected: %s, got: %s", want, string(json))
	}
}

func TestUserWithToken(t *testing.T) {
	want := "{\"email\":\"abc@example.com\",\"password\":\"awwwww\",\"token\":{\"String\":\"abc\",\"Valid\":true}}"
	u := User{Email: "abc@example.com", Password: "awwwww", Token: sql.NullString{String: "abc", Valid: true}}

	json, err := json.Marshal(u)
	if err != nil {
		t.Errorf("json marshal error: %s", err)
	}
	if string(json) != want {
		t.Errorf("Marshaled json expected: %s, got: %s", want, string(json))
	}
}
