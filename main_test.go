package main

import (
	"fmt"
	"testing"
)

func TestSystemRunning(t *testing.T) {
	message := "System is running in the port "
	defaultPort := 8000

	t.Run("System running in the port", func(t *testing.T) {
		const port int = 3000

		got := SystemRunningMessage(port)
		want := message + fmt.Sprint(port)

		assertCorrectMessage(t, got, want)
	})

	t.Run("System running without a port", func(t *testing.T) {
		const port int = 0

		got := SystemRunningMessage(port)
		want := message + fmt.Sprint(defaultPort)

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		const errMessage string = "got %q want %q"
		t.Errorf(errMessage, got, want)
	}
}
