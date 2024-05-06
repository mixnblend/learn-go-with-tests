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
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 12.0, Height: 6.00}, 72.0},
		{Circle{Radius: 10}, 314.1592653589793},
		{Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.want
		if got != want {
			t.Errorf("wanted %g, got %g", want, got)
		}
	}
}
