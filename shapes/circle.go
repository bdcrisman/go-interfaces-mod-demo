package shapes

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return float64(math.Pi * math.Pow(c.Radius, 2))
}

func (c Circle) Volume() float64 {
	return float64(4 / 3 * math.Pi * math.Pow(float64(c.Radius), 3))
}
