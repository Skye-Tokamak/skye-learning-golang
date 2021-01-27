package myinterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 4.0}
	got := Perimeter(rectangle)
	want := 28.0
	if got != want {
		t.Errorf("got '%.2f' but want '%.2f'", got, want)
	}
}
func TestAria(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got '%.2f' but want '%.2f'", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 4.0}
		checkArea(t, rectangle, 40.0)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
