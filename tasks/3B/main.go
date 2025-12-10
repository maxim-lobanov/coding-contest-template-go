package main

import (
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/algo"
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	result := 0
	for i := 0; i < len(input); i++ {
		bank := algo.Map(strings.Split(input[i], ""), func(s string) int { return cast.ParseInt(s) })
		result += solveLine(bank)
	}
	return cast.ToString(result)
}

func solveLine(bank []int) int {
	result := 0
	windowSize := 12
	lastFoundIndex := -1
	for windowSize > 0 {
		foundIndex := lastFoundIndex + 1
		for i := lastFoundIndex + 1; i <= len(bank)-windowSize; i++ {
			if bank[i] > bank[foundIndex] {
				foundIndex = i
			}
		}

		result = result*10 + bank[foundIndex]
		lastFoundIndex = foundIndex
		windowSize--
	}

	return result
}
