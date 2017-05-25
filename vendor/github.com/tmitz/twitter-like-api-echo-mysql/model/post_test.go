package model

import (
	"encoding/json"
	"testing"
)

func TestPost(t *testing.T) {
	want := "{\"receive_id\":111,\"send_id\":222,\"message\":\"hello,world!\"}"
	p := Post{ReceiveID: 111, SendID: 222, Message: "hello,world!"}
	json, err := json.Marshal(p)
	if err != nil {
		t.Errorf("json marshal error: %s", err)
	}
	if string(json) != want {
		t.Errorf("Marshaled json expected: %s, got: %s", want, string(json))
	}
}
