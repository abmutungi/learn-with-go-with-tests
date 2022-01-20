package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Arnold")
	want := "Hello, Arnold"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
