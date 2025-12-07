package main

import (
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	allValues := [][]int{}
	for i := 0; i < len(input)-1; i++ {
		numsRow := cast.ParseIntArray(input[i])
		for len(allValues) < len(numsRow) {
			allValues = append(allValues, []int{})
		}
		for j := 0; j < len(numsRow); j++ {
			allValues[j] = append(allValues[j], numsRow[j])
		}
	}

	result := 0
	operations := cast.ParseStringArray(input[len(input)-1])
	for opInd, op := range operations {
		opValues := allValues[opInd]
		acc := opValues[0]

		for j := 1; j < len(opValues); j++ {
			switch op {
			case "+":
				acc += opValues[j]
			case "*":
				acc *= opValues[j]
			default:
				panic("unknown operation")
			}
		}

		result += acc
	}

	return cast.ToString(result)
}
