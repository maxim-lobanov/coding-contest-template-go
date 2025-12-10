package main

import (
	"fmt"
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/algo"
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	totalResult := 0
	for _, line := range input {
		result := solveSingleCase(line)
		totalResult += result
	}
	return cast.ToString(totalResult)
}

func solveSingleCase(input string) int {
	requiredPattern, availableOptions := parseinputLine(input)

	marked := map[string]int{}
	queue := algo.Queue[string]{}

	startPattern := strings.ReplaceAll(requiredPattern, "#", ".")

	queue.PushBack(startPattern)
	marked[startPattern] = 0

	for marked[requiredPattern] == 0 {
		if queue.Count() == 0 {
			panic("queue is empty but solution is not found")
		}

		currentPattern, _ := queue.TakeFront()
		for _, optionToApply := range availableOptions {
			newPattern := applyOptionToPattern(currentPattern, optionToApply)
			if marked[newPattern] == 0 {
				marked[newPattern] = marked[currentPattern] + 1
				queue.PushBack(newPattern)
			}
		}
	}

	return marked[requiredPattern]
}

func applyOptionToPattern(pattern string, option []int) string {
	runes := []rune(pattern)
	for _, optionItem := range option {
		if runes[optionItem] == '#' {
			runes[optionItem] = '.'
		} else if runes[optionItem] == '.' {
			runes[optionItem] = '#'
		} else {
			panic(fmt.Sprintf("invalid pattern: %s", pattern))
		}
	}

	return string(runes)
}

func parseinputLine(input string) (string, [][]int) {
	if !strings.HasPrefix(input, "[") || !strings.HasSuffix(input, "}") {
		panic(fmt.Sprintf("invalid input line: %s", input))
	}

	parts := strings.Split(input, "]")
	requiredPattern := strings.TrimPrefix(parts[0], "[")

	parts = strings.Split(parts[1], "{")
	optionsRaw := strings.Split(strings.TrimSpace(parts[0]), " ")

	availableOptions := [][]int{}
	for _, optionRawPart := range optionsRaw {
		optionRawPart = strings.TrimSpace(optionRawPart)
		if !strings.HasPrefix(optionRawPart, "(") || !strings.HasSuffix(optionRawPart, ")") {
			panic(fmt.Sprintf("invalid option: %s", optionRawPart))
		}

		optionRawPart = strings.TrimPrefix(optionRawPart, "(")
		optionRawPart = strings.TrimSuffix(optionRawPart, ")")
		optionRawList := strings.Split(optionRawPart, ",")
		optionsList := algo.Map(optionRawList, func(item string) int { return cast.ParseInt(item) })
		availableOptions = append(availableOptions, optionsList)
	}

	return requiredPattern, availableOptions
}
