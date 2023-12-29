package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	var width, height float64 = 10.0, 10.0

	got := Perimeter(width, height)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestRectangle(t *testing.T) {
	var width, height float64 = 12.0, 6.0

	got := Rectangle(width, height)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
