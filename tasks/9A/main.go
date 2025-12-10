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
			if square > result {
				result = square
			}
		}
	}

	return cast.ToString(result)
}
