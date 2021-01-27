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
	//表格驱动测试
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rect", shape: Rectangle{Width: 10.0, Height: 4.0}, hasArea: 40.0},
		{name: "sphere", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "tri", shape: Triangle{Base: 10.0, Height: 10.0}, hasArea: 50.0},
	}
	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%v got %.2f want %.2f", tt.shape, got, tt.hasArea)
		}
	}

}
