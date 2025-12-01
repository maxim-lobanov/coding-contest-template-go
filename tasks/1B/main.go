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
		sign := 1
		if dir == 'L' {
			sign = -1
		}

		if steps >= 100 {
			result += steps / 100
			steps = steps % 100
		}
		newDialValue := dialValue + sign*steps
		if newDialValue < 0 {
			newDialValue += 100
			if dialValue > 0 {
				result++
			}
		} else if newDialValue >= 100 {
			newDialValue -= 100
			result++
		} else if newDialValue == 0 {
			result++
		}
		dialValue = newDialValue
	}

	return cast.ToString(result)
}
