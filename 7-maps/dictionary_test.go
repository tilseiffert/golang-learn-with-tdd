package maps

import "testing"

func TestSearch(t *testing.T) {
	dicitonary := map[string]string{"test": "this is just a test"}

	searchkey := "test"
	got := Search(dicitonary, searchkey)
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, searchkey)
	}
}
