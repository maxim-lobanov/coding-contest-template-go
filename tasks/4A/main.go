package main

import (
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	result := 0

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if checkPoint(input, x, y) {
				result++
			}
		}
	}

	return cast.ToString(result)
}

func checkPoint(matrix []string, x, y int) bool {
	const symbol = '@'
	if matrix[y][x] != symbol {
		return false
	}

	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	foundAround := 0
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if newX >= 0 && newX < len(matrix[0]) && newY >= 0 && newY < len(matrix) {
			if matrix[newY][newX] == symbol {
				foundAround++
			}
		}
	}

	return foundAround < 4
}
