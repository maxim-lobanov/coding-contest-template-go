package main

import (
	"fmt"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	lineIndex := 0
	ranges := make([]Range, 0)
	for input[lineIndex] != "" {
		ranges = append(ranges, parseRange(input[lineIndex]))
		lineIndex++
	}

	result := 0
	for i := lineIndex + 1; i < len(input); i++ {
		value := cast.ParseInt(input[i])
		if isIncludedToAnyRange(ranges, value) {
			result++
		}
	}

	return cast.ToString(result)
}

func parseRange(line string) Range {
	r := Range{}
	fmt.Sscanf(line, "%d-%d", &r.Start, &r.End)
	return r
}

func isIncludedToAnyRange(ranges []Range, value int) bool {
	for _, r := range ranges {
		if r.Includes(value) {
			return true
		}
	}
	return false
}

type Range struct {
	Start, End int
}

func (r Range) Includes(value int) bool {
	return value >= r.Start && value <= r.End
}
