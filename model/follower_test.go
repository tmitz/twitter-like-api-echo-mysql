package model

import (
	"encoding/json"
	"testing"
)

func TestFollower(t *testing.T) {
	want := "{\"follower_id\":999,\"followed_id\":888}"
	f := Follower{999, 888}
	json, err := json.Marshal(f)
	if err != nil {
		t.Errorf("json marshal error: %s", err)
	}
	if string(json) != want {
		t.Errorf("Marshaled json expected: %s, got: %s", want, string(json))
	}
}
