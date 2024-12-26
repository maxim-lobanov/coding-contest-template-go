package main

import (
	"fmt"
	"strconv"

	taskRunner "github.com/maxim-lobanov/coding-contest-template-go/runner"
)

func main() {
	// runner.Test(runner.InputFromText("1\n2\n3"), runner.OutputFromText("6"))
	// runner.Test(runner.InputFromFile("test1.in"), runner.OutputFromFile("test1.out"))

	runner := taskRunner.NewRunner(run)

	runner.Test(runner.InputFromText("1\n2\n3"), runner.OutputFromText("6"))
	runner.Main(runner.InputFromFile("main.in"))

	runner.Execute()
}

func run(input []string) string {
	num1, _ := strconv.Atoi(input[0])
	num2, _ := strconv.Atoi(input[1])
	num3, _ := strconv.Atoi(input[2])

	result := num1 + num2 + num3
	return fmt.Sprint(result)
}
