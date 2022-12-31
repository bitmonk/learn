package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rect Rectangle) (perim float64) {
	perim = 2 * (rect.Width + rect.Height)
	return
}

func (r Rectangle) Area() (area float64) {
	area = r.Width * r.Height
	return
}

func (c Circle) Area() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return
}

func (t Triangle) Area() (area float64) {
	area = (t.Base * t.Height) * 0.5
	return
}
