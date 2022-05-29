package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Tilmann", "")
		want := "Hello, Tilmann!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Emil", "German")
		want := "Hallo, Emil!"
		assertCorrectMessage(t, got, want)
	})

	// t.Run("in French", func(t *testing.T) {
	// 	got := Hello("Julius", "French")
	// 	want := "Bonjour, Julius‚!"
	// 	assertCorrectMessage(t, got, want)
	// })

}
