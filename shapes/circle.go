package shapes

import "math"

// Circle type
type Circle struct {
	Radius float64
}

// Area is implemented from ishape
func (c Circle) Area() float64 {
	return float64(math.Pi * math.Pow(c.Radius, 2))
}

// Volume is implemented from ishape
func (c Circle) Volume() float64 {
	return float64(4 / 3 * math.Pi * math.Pow(float64(c.Radius), 3))
}
