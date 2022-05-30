package geometry

import "testing"

func checkFloats(t testing.TB, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	checkFloats(t, got, want)

}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0
		checkFloats(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		checkFloats(t, got, want)
	})

}
