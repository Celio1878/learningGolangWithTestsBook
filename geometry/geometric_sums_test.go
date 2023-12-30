package geometry

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", Measures{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}
	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.hasArea)
		})
	}
}

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()

	got := shape.Area()

	if got != want {
		t.Errorf("%#v got %g want %g", shape, got, want)
	}
}
