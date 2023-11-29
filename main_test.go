package main

import "testing"

func TestSystemRunning(t *testing.T) {
	got := SystemRunningMessage()
	want := "System is running"

	if got != want {
		const errMessage string = "got %q want %q"
		t.Errorf(errMessage, got, want)
	}

}
