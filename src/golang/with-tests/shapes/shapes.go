package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rect Rectangle) (perim float64) {
	perim = 2 * (rect.Width + rect.Height)
	return
}

func Area(rect Rectangle) (area float64) {
	area = rect.Width * rect.Height
	return
}
