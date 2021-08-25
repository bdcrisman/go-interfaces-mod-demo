package shapes

type Square struct {
	Length, Width, Height float64
}

func (s Square) Area() float64 {
	return float64(s.Length * s.Width)
}

func (s Square) Volume() float64 {
	return float64(s.Length * s.Width * s.Height)
}
