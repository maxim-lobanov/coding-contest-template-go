package main

import (
	"fmt"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	num1 := cast.ParseInt(input[0])
	num2 := cast.ParseInt(input[1])
	num3 := cast.ParseInt(input[2])
	p := 0
	if num1 != 10 {
		p = 1
	}

	result := num1 + num2 + num3/p
	return fmt.Sprint(result)
}
