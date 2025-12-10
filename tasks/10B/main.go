package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/algo"
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	totalResult := 0
	for index, line := range input {
		fmt.Printf("%d/%d\n", index, len(input))
		result := solveSingleCase(line)
		totalResult += result
	}
	return cast.ToString(totalResult)
}

func solveSingleCase(input string) int {
	requiredPattern, availableButtons := parseinputLine(input)

	numVars := len(availableButtons)
	equations := make([][]int, len(requiredPattern))
	for i := 0; i < len(requiredPattern); i++ {
		equations[i] = make([]int, numVars+1)
		for j := 0; j < numVars; j++ {
			if slices.Contains(availableButtons[j], i) {
				equations[i][j] = 1
			}
		}
		equations[i][numVars] = requiredPattern[i]
	}

	result, found := solveLinearSystem(equations, numVars)
	if !found {
		panic("solution is not found")
	}

	return result
}

func parseinputLine(input string) ([]int, [][]int) {
	if !strings.HasPrefix(input, "[") || !strings.HasSuffix(input, "}") {
		panic(fmt.Sprintf("invalid input line: %s", input))
	}

	parts := strings.Split(input, "]")

	parts = strings.Split(parts[1], "{")
	buttonsRaw := strings.Split(strings.TrimSpace(parts[0]), " ")

	requiredPatternRaw := strings.TrimSuffix(parts[1], "}")
	requiredPatternParts := strings.Split(requiredPatternRaw, ",")
	requiredPattern := algo.Map(requiredPatternParts, func(item string) int { return cast.ParseInt(item) })

	availableButtons := [][]int{}
	for _, buttonRawPart := range buttonsRaw {
		buttonRawPart = strings.TrimSpace(buttonRawPart)
		if !strings.HasPrefix(buttonRawPart, "(") || !strings.HasSuffix(buttonRawPart, ")") {
			panic(fmt.Sprintf("invalid option: %s", buttonRawPart))
		}

		buttonRawPart = strings.TrimPrefix(buttonRawPart, "(")
		buttonRawPart = strings.TrimSuffix(buttonRawPart, ")")
		optionRawList := strings.Split(buttonRawPart, ",")
		optionsList := algo.Map(optionRawList, func(item string) int { return cast.ParseInt(item) })
		availableButtons = append(availableButtons, optionsList)
	}

	return requiredPattern, availableButtons
}

func solveLinearSystem(equations [][]int, numVars int) (int, bool) {
	coefficients := make([][]int, len(equations))
	rightSide := make([]int, len(equations))

	for i, eq := range equations {
		coefficients[i] = eq[:numVars]
		rightSide[i] = eq[numVars]
	}

	current := make([]int, numVars)
	currentSums := make([]int, len(equations))
	minSum := math.MaxInt
	found := false

	var backtrack func(varIndex, currentSum int)
	backtrack = func(varIndex, currentSum int) {
		if currentSum >= minSum {
			return
		}

		if varIndex == numVars {
			valid := true
			for i := 0; i < len(equations); i++ {
				if currentSums[i] != rightSide[i] {
					valid = false
					break
				}
			}
			if valid {
				found = true
				if currentSum < minSum {
					minSum = currentSum
				}
			}
			return
		}

		maxVal := getMaxValue(varIndex, currentSums, coefficients, rightSide)
		if maxVal < 0 {
			return
		}

		lowerBound := 0
		for i := 0; i < len(equations); i++ {
			remaining := rightSide[i] - currentSums[i]
			if remaining > lowerBound {
				lowerBound = remaining
			}
		}
		if varIndex < numVars && lowerBound > 0 {
			lowerBound = lowerBound / (numVars - varIndex)
		}
		if currentSum+lowerBound >= minSum {
			return
		}

		for val := 0; val <= maxVal; val++ {
			current[varIndex] = val
			for i := 0; i < len(equations); i++ {
				currentSums[i] += coefficients[i][varIndex] * val
			}

			backtrack(varIndex+1, currentSum+val)

			for i := 0; i < len(equations); i++ {
				currentSums[i] -= coefficients[i][varIndex] * val
			}
		}
		current[varIndex] = 0
	}

	backtrack(0, 0)

	if !found {
		return 0, false
	}
	return minSum, true
}

func getMaxValue(varIndex int, currentSums []int, coefficients [][]int, rightSide []int) int {
	maxVal := math.MaxInt
	hasConstraint := false

	for i := 0; i < len(coefficients); i++ {
		if coefficients[i][varIndex] > 0 {
			remaining := rightSide[i] - currentSums[i]
			if remaining < 0 {
				return -1
			}
			possible := remaining / coefficients[i][varIndex]
			if possible < maxVal {
				maxVal = possible
				hasConstraint = true
			}
		}
	}

	if !hasConstraint {
		maxVal = 0
		for i := 0; i < len(rightSide); i++ {
			maxVal += rightSide[i]
		}
	}

	return maxVal
}
