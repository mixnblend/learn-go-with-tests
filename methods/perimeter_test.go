package methods

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("wanted %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(shape Shape, want float64, t testing.TB) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("wanted %g, got %g", want, got)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.00}
		want := 72.00
		checkArea(rectangle, want, t)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(circle, want, t)
	})
}
