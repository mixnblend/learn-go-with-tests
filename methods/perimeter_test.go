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
	rectangle := Rectangle{12.0, 6.00}
	got := Area(rectangle)
	want := 72.00

	if got != want {
		t.Errorf("wanted %.2f, got %.2f", want, got)
	}
}
