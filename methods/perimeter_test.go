package methods

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("wanted %.2f, got %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.00

	if got != want {
		t.Errorf("wanted %.2f, got %.2f", got, want)
	}
}
