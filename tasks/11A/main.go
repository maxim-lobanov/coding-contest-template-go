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
	waysCount := map[string]int{}

	queue.PushBack("you")
	waysCount["you"] = 1
	inputsCount["you"] = 0

	for queue.Count() > 0 {
		currentItem, _ := queue.TakeFront()

		if inputsCount[currentItem] > 0 {
			panic("current item has unsivited inputs")
		}

		if waysCount[currentItem] == 0 {
			panic("current item has 0 ways")
		}

		for _, toItem := range connections[currentItem] {
			if inputsCount[toItem] < 1 {
				panic("toItem has 0 remain inputs")
			}

			waysCount[toItem] += waysCount[currentItem]

			inputsCount[toItem]--
			if inputsCount[toItem] == 0 {
				queue.PushBack(toItem)
			}
		}
	}

	result := waysCount["out"]
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
	queue.PushBack("you")
	visited["you"] = true

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
