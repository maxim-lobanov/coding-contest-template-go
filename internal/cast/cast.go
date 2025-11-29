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
	result := make([]int, len(parts))
	for i, part := range parts {
		result[i] = ParseInt(part)
	}

	return result
}
