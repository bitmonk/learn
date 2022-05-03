package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
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
	area = 0
	return
}
