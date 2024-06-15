package main

import "math"

type Cube struct {
	Side float64
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}

func (c Cube) Wide() float64 {
	return math.Pow(c.Side, 2) * 6
}

func (c Cube) Circular() float64 {
	return c.Side * 12
}
