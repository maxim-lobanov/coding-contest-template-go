package main

import (
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	allValues := parseInput(input[:len(input)-1])

	result := 0
	operations := cast.ParseStringArray(input[len(input)-1])
	for opInd, op := range operations {
		opValues := convertValuesToNewNotation(allValues[opInd])
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

func parseInput(inputLines []string) [][]string {
	maxLineLength := 0
	for _, inputLine := range inputLines {
		maxLineLength = max(maxLineLength, len(inputLine))
	}

	isEmptyColumn := func(colIndex int) bool {
		for _, inputLine := range inputLines {
			if colIndex < len(inputLine) && inputLine[colIndex] != ' ' {
				return false
			}
		}

		return true
	}
	getValuesFromAllLines := func(colStart, colEnd int) []string {
		vals := make([]string, len(inputLines))

		for ind, line := range inputLines {
			vals[ind] = line[colStart:min(colEnd, len(line))]
		}

		return vals
	}

	emptyColumns := []int{}
	for i := 0; i < maxLineLength; i++ {
		if isEmptyColumn(i) {
			emptyColumns = append(emptyColumns, i)
		}
	}

	parsedValues := make([][]string, len(emptyColumns)+1)
	for i := 0; i < len(emptyColumns); i++ {
		colStart := 0
		if i > 0 {
			colStart = emptyColumns[i-1] + 1
		}
		colEnd := emptyColumns[i]
		parsedValues[i] = getValuesFromAllLines(colStart, colEnd)
	}
	parsedValues[len(emptyColumns)] = getValuesFromAllLines(emptyColumns[len(emptyColumns)-1]+1, maxLineLength)

	return parsedValues
}

func convertValuesToNewNotation(originalValues []string) []int {
	maxValueLen := 0
	for _, val := range originalValues {
		maxValueLen = max(maxValueLen, len(val))
	}

	newValues := make([]int, maxValueLen)
	for i := 0; i < maxValueLen; i++ {
		newValStr := ""
		for _, orValue := range originalValues {
			if orValue[i] != ' ' {
				newValStr += string(orValue[i])
			}
		}
		newValParsed := cast.ParseInt(newValStr)
		newValues[i] = newValParsed
	}

	return newValues
}
