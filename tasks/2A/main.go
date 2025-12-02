package main

import (
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	ranges := strings.Split(input[0], ",")
	totalResult := 0

	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		left := cast.ParseInt(bounds[0])
		right := cast.ParseInt(bounds[1])
		result := solveRange(left, right)

		totalResult += result
	}

	return cast.ToString(totalResult)
}

func solveRange(left, right int) int {
	result := 0
	for num := left; num <= right; num++ {
		if isInvalidNumber(cast.ToString(num)) {
			result += num
		}
	}

	return result
}

func isInvalidNumber(num string) bool {
	if len(num)%2 != 0 {
		return false
	}

	return isRepeatedPattern(num, len(num)/2)
}

func isRepeatedPattern(num string, patternLen int) bool {
	if len(num)%patternLen != 0 {
		return false
	}

	pattern := num[:patternLen]
	for i := patternLen; i < len(num); i += patternLen {
		if num[i:i+patternLen] != pattern {
			return false
		}
	}

	return true
}
