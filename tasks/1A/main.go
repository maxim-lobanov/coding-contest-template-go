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
		switch dir {
		case 'L':
			dialValue -= steps
		case 'R':
			dialValue += steps
		}

		dialValue = normalizeDialValue(dialValue)
		if dialValue == 0 {
			result++
		}
	}

	return cast.ToString(result)
}

func normalizeDialValue(value int) int {
	for value < 0 {
		value += 100
	}

	for value >= 100 {
		value -= 100
	}

	return value
}
