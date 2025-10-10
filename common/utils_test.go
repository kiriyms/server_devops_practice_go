package common

import (
	"strings"
	"testing"
)

func TestGetUserId(t *testing.T) {
	id1 := GetUserId()
	id2 := GetUserId()

	if id1 == "" || id2 == "" {
		t.Fatal("expected non-empty user IDs")
	}
	if id1 == id2 {
		t.Fatal("expected unique user IDs")
	}
	if !strings.Contains(id1, "_") {
		t.Errorf("expected underscore in user ID, got %s", id1)
	}
}
