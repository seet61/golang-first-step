package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := HelloN(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
