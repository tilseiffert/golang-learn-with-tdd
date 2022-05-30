package geometry

import "testing"

func checkFloats(t testing.TB, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {

	//t.Run("check Perimeter", func(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	checkFloats(t, got, want)
	//})
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0
	checkFloats(t, got, want)
}
