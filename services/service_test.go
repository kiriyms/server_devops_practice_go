package services

import (
	"context"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	g := NewGreeter()
	msg, err := g.Greet(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.HasPrefix(msg, "Hallo, user ") {
		t.Errorf("unexpected greeting format: %s", msg)
	}

	if !strings.Contains(msg, "_") {
		t.Errorf("expected user ID with underscore in greeting, got: %s", msg)
	}
}
