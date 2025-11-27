package main

import (
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
	"github.com/maxim-lobanov/coding-contest-template-go/internal/runner"
)

func solution(input []string, output *runner.OutputBuilder) {
	num1 := cast.ParseInt(input[0])
	num2 := cast.ParseInt(input[1])
	num3 := cast.ParseInt(input[2])

	result := num1 + num2 + num3
	output.WriteInt(result)
}
