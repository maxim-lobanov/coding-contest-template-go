package main

import (
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/algo"
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	points := make([]algo.Point, len(input))
	for i := 0; i < len(input); i++ {
		parts := strings.Split(input[i], ",")
		points[i] = algo.Point{
			X: cast.ParseInt(parts[0]),
			Y: cast.ParseInt(parts[1]),
		}
	}

	result := -1
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			square := (algo.Abs(p1.X-p2.X) + 1) * (algo.Abs(p1.Y-p2.Y) + 1)
			if square > result && isRectangleFullyInsidePolygon(p1, p2, points) {
				result = square
			}
		}
	}

	return cast.ToString(result)
}

func isRectangleFullyInsidePolygon(pA, pB algo.Point, polygon []algo.Point) bool {
	corners := [4]algo.Point{
		{X: min(pA.X, pB.X), Y: min(pA.Y, pB.Y)},
		{X: max(pA.X, pB.X), Y: min(pA.Y, pB.Y)},
		{X: max(pA.X, pB.X), Y: max(pA.Y, pB.Y)},
		{X: min(pA.X, pB.X), Y: max(pA.Y, pB.Y)},
	}

	// Check that all corners are inside or on the boundary of polygon
	for _, corner := range corners {
		if !isPointInOrOnPolygon(corner, polygon) {
			return false
		}
	}

	// Check that no polygon edges cross through the rectangle interior
	rectEdges := [][2]algo.Point{
		{corners[0], corners[1]},
		{corners[1], corners[2]},
		{corners[2], corners[3]},
		{corners[3], corners[0]},
	}

	for i := 0; i < len(polygon); i++ {
		j := (i + 1) % len(polygon)

		for _, rectEdge := range rectEdges {
			if doSegmentsCrossIntersect(polygon[i], polygon[j], rectEdge[0], rectEdge[1]) {
				return false
			}
		}
	}

	return true
}

func isPointInOrOnPolygon(p algo.Point, polygon []algo.Point) bool {
	// First check if point is a polygon vertex
	for _, vertex := range polygon {
		if p.X == vertex.X && p.Y == vertex.Y {
			return true
		}
	}

	// Then check if inside using ray casting
	polygonSize := len(polygon)
	inside := false

	j := polygonSize - 1
	for i := 0; i < polygonSize; i++ {
		xi, yi := polygon[i].X, polygon[i].Y
		xj, yj := polygon[j].X, polygon[j].Y

		if ((yi > p.Y) != (yj > p.Y)) && (p.X < (xj-xi)*(p.Y-yi)/(yj-yi)+xi) {
			inside = !inside
		}
		j = i
	}

	return inside
}

func doSegmentsCrossIntersect(p1, p2, p3, p4 algo.Point) bool {
	// Only return true if segments properly intersect (not just touch at endpoints)
	d1 := direction(p3, p4, p1)
	d2 := direction(p3, p4, p2)
	d3 := direction(p1, p2, p3)
	d4 := direction(p1, p2, p4)

	// Proper intersection (segments cross)
	if ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) &&
		((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0)) {
		return true
	}

	return false
}

func direction(p1, p2, p3 algo.Point) int {
	// Cross product
	val := (p3.Y-p1.Y)*(p2.X-p1.X) - (p2.Y-p1.Y)*(p3.X-p1.X)
	if val == 0 {
		return 0 // Collinear
	}
	if val > 0 {
		return 1 // Clockwise
	}
	return -1 // Counterclockwise
}
