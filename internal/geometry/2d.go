package geometry

import "math"

type Point struct {
	X int
	Y int
}

func (p Point) DistanceToPoint(other Point) float64 {
	dx := float64(other.X - p.X)
	dy := float64(other.Y - p.Y)
	return math.Sqrt(dx*dx + dy*dy)
}
