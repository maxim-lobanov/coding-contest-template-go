package main

import (
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/algo"
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	connections := parseInput(input)
	inputsCount := getInputsCount(connections)

	queue := algo.Queue[string]{}
	queue.PushBack("svr")
	inputsCount["svr"] = 0

	waysCountNone := map[string]int{"svr": 1} // the count of ways to come to X without visiting "dac" or "fft" vertexes
	waysCountDac := map[string]int{"svr": 0}  // the count of ways to come to X with visiting only "dac" vextex
	waysCountFft := map[string]int{"svr": 0}  // the count of ways to come to X with visiting only "fft" vertex
	waysCountBoth := map[string]int{"svr": 0} // the count of ways to come to X with visiting both "dac" and "fft"

	for queue.Count() > 0 {
		currentItem, _ := queue.TakeFront()

		if inputsCount[currentItem] > 0 {
			panic("current item has unsivited inputs")
		}

		if waysCountNone[currentItem] == 0 {
			panic("current item has 0 ways")
		}

		for _, toItem := range connections[currentItem] {
			if inputsCount[toItem] < 1 {
				panic("toItem has 0 remain inputs")
			}

			if toItem == "dac" {
				waysCountNone[toItem] += waysCountNone[currentItem]
				waysCountDac[toItem] += waysCountNone[currentItem] + waysCountDac[currentItem]
				waysCountFft[toItem] += 0 // No paths can have only fft after visiting dac
				waysCountBoth[toItem] += waysCountFft[currentItem] + waysCountBoth[currentItem]
			} else if toItem == "fft" {
				waysCountNone[toItem] += waysCountNone[currentItem]
				waysCountDac[toItem] += 0 // No paths can have only dac after visiting fft
				waysCountFft[toItem] += waysCountNone[currentItem] + waysCountFft[currentItem]
				waysCountBoth[toItem] += waysCountDac[currentItem] + waysCountBoth[currentItem]
			} else {
				waysCountNone[toItem] += waysCountNone[currentItem]
				waysCountDac[toItem] += waysCountDac[currentItem]
				waysCountFft[toItem] += waysCountFft[currentItem]
				waysCountBoth[toItem] += waysCountBoth[currentItem]
			}

			inputsCount[toItem]--
			if inputsCount[toItem] == 0 {
				queue.PushBack(toItem)
			}
		}
	}

	result := waysCountBoth["out"]
	return cast.ToString(result)
}

func parseInput(input []string) map[string][]string {
	connections := map[string][]string{}

	for _, line := range input {
		parts := strings.Split(line, ":")
		from := strings.TrimSpace(parts[0])
		toList := strings.Split(parts[1], " ")
		for _, toItem := range toList {
			toItem = strings.TrimSpace(toItem)
			if toItem == "" {
				continue
			}

			connections[from] = append(connections[from], toItem)
		}
	}

	return connections
}

func getInputsCount(connections map[string][]string) map[string]int {
	visited := map[string]bool{}

	queue := algo.Queue[string]{}
	queue.PushBack("svr")
	visited["svr"] = true

	for queue.Count() > 0 {
		currentItem, _ := queue.TakeFront()

		for _, toItem := range connections[currentItem] {
			if found := visited[toItem]; !found {
				queue.PushBack(toItem)
				visited[toItem] = true
			}
		}
	}

	inputsCount := map[string]int{}
	for from, toList := range connections {
		if found := visited[from]; !found {
			continue
		}

		for _, toItem := range toList {
			inputsCount[toItem]++
		}
	}

	return inputsCount
}
