package main

import (
	"testing"
)

func TestTestePrint(t *testing.T) {
	want := "Por enquanto..."
	if got := TestePrint(); got != want {
		t.Errorf("TestePrint() = %s, want %q", got, want)
	}
}
