package geometry

import (
	"strconv"
	"testing"
)

func checkFloats(t testing.TB, got float64, want float64, hint string) {

	t.Helper()

	if hint != "" {
		hint = ", " + hint
	}

	if got != want {
		t.Errorf("got %g want %g%s", got, want, hint)
	}

}

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	checkFloats(t, got, want, "")

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		testname string
		shape    Shape
		want     float64
	}{
		{"rectangle", Rectangle{12, 6}, 72.0},
		{"circle", Circle{10}, 314.1592653589793},
		{"triangle", Triangle{12, 6}, 36.0},
	}

	for i, tt := range areaTests {
		got := tt.shape.Area()
		checkFloats(t, got, tt.want, tt.testname+" (i="+strconv.Itoa(i)+")")
	}

}
