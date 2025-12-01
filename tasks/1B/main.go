package main

import (
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	dialValue := 50
	result := 0

	for _, line := range input {
		dir := line[0]
		steps := cast.ParseInt(line[1:])
		side := 1
		if dir == 'L' {
			side = -1
		}

		newDialValue, rotations := rotate(dialValue, side, steps)
		dialValue = newDialValue
		result += rotations
	}

	return cast.ToString(result)
}

func rotate(current, side, steps int) (int, int) {
	localResult := 0

	for i := 0; i < steps; i++ {
		current += side
		if current < 0 {
			current = 99
		} else if current >= 100 {
			current = 0
			localResult++
		} else if current == 0 {
			localResult++
		}
	}

	return current, localResult
}
