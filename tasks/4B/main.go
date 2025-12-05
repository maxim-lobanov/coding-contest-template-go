package main

import (
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	result := 0
	matrix := readInput(input)

	for {
		newMatrix, changed := runCleanupIteration(matrix)
		if changed > 0 {
			result += changed
			matrix = newMatrix
		} else {
			break
		}
	}

	return cast.ToString(result)
}

func runCleanupIteration(matrix [][]rune) ([][]rune, int) {
	changed := 0
	newMatrix := make([][]rune, len(matrix))
	for y := 0; y < len(matrix); y++ {
		newRow := make([]rune, len(matrix[y]))
		for x := 0; x < len(matrix[y]); x++ {
			if checkPoint(matrix, x, y) {
				newRow[x] = '.'
				changed++
			} else {
				newRow[x] = matrix[y][x]
			}
		}
		newMatrix[y] = newRow
	}

	return newMatrix, changed
}

func checkPoint(matrix [][]rune, x, y int) bool {
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

func readInput(input []string) [][]rune {
	matrix := make([][]rune, len(input))
	for i := range input {
		matrix[i] = make([]rune, len(input[i]))
		for j := range input[i] {
			matrix[i][j] = rune(input[i][j])
		}
	}

	return matrix
}
