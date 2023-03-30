package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "RectangleArea", shape: Rectangle{12, 6}, want: 72.0},
		{name: "CircleArea", shape: Circle{10}, want: 314.1592653589793},
		{name: "TriangleArea", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}
