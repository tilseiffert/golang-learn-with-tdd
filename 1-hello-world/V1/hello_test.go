package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tilmann")
	want := "Hello, Tilmann!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
