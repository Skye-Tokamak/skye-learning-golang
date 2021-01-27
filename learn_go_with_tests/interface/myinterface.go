package myinterface

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}
type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	peri := 2 * (rectangle.Width + rectangle.Height)
	return peri
}
func Aria(rectangle Rectangle) float64 {
	aria := rectangle.Width * rectangle.Height
	return aria
}
