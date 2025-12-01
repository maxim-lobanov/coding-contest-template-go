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

		dialValue %= 100
		if dialValue < 0 {
			dialValue += 100
		} else if dialValue >= 100 {
			dialValue -= 100
		}
		if dialValue == 0 {
			result++
		}
	}

	return cast.ToString(result)
}
