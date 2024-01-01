package greetings

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Someone")

	got := buffer.String()
	want := "Hello, Someone!\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
