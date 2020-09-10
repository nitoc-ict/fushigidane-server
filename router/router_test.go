package router

import (
	"testing"
)

func TestNewRouter(t *testing.T) {
	res := NewRouter()

	if res != "Hello" {
		t.Fatal("failed test")
	}
}
