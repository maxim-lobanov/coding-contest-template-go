package algo

import "math"

type Point3D struct {
	X int
	Y int
	Z int
}

func (p Point3D) DistanceToPoint(other Point3D) float64 {
	dx := float64(other.X - p.X)
	dy := float64(other.Y - p.Y)
	dz := float64(other.Z - p.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
