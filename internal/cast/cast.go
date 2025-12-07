package cast

import (
	"fmt"
	"strconv"
	"strings"
)

func ToPtr[T any](value T) *T {
	return &value
}

func ToString[T any](value T) string {
	return fmt.Sprintf("%v", value)
}

func ParseInt(s string) int {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(res)
}

func ParseInt64(s string) int64 {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func ParseFloat(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func ParseIntArray(s string) []int {
	parts := strings.Split(s, " ")
	result := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, ParseInt(part))
		}
	}

	return result
}

func ParseStringArray(s string) []string {
	parts := strings.Split(s, " ")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}

func ParseCharMatrix(lines []string) [][]rune {
	result := make([][]rune, len(lines))
	for i, line := range lines {
		result[i] = []rune(line)
	}

	return result
}
