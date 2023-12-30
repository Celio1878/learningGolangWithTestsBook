package geometry

import "math"

type Shape interface {
	Area() float64
}

type Measures struct {
	Width, Height float64
}

func (m Measures) Area() float64 {
	return m.Width * m.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
