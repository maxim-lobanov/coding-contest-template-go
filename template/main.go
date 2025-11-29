package main

import (
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	num1 := cast.ParseInt(input[0])
	num2 := cast.ParseInt(input[1])
	num3 := cast.ParseInt(input[2])

	result := num1 + num2 + num3
	return cast.ToString(result)
}
