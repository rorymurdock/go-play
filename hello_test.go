package main

import "testing"

func Test_Hello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Test_Bye(t *testing.T) {
	got := Bye()
	want := "Bye!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
