package main

import (
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
	"github.com/maxim-lobanov/coding-contest-template-go/internal/linq"
)

func solution(input []string) string {
	result := 0
	for i := 0; i < len(input); i++ {
		bank := linq.Map(strings.Split(input[i], ""), func(s string) int { return cast.ParseInt(s) })
		result += solveLine(bank)
	}
	return cast.ToString(result)
}

func solveLine(bank []int) int {
	firstMaxIndex := 0
	for i := 1; i < len(bank)-1; i++ {
		if bank[i] > bank[firstMaxIndex] {
			firstMaxIndex = i
		}
	}

	secondMaxIndex := firstMaxIndex + 1
	for i := secondMaxIndex + 1; i < len(bank); i++ {
		if bank[i] > bank[secondMaxIndex] {
			secondMaxIndex = i
		}
	}

	return bank[firstMaxIndex]*10 + bank[secondMaxIndex]
}
