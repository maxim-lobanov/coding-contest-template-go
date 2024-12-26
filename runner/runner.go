package runner

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type RunFunc = func(input []string) string

type Runner struct {
	runFunc           RunFunc
	taskRootDirectory string
	testCases         []*runCase
	mainCase          *runCase
}

func NewRunner(runFunc RunFunc) *Runner {
	return &Runner{
		runFunc:           runFunc,
		taskRootDirectory: resolveTaskRootDirectory(),
		testCases:         []*runCase{},
		mainCase:          nil,
	}
}

func (r *Runner) Execute() {
	testsResult := true
	for testCaseIndex, _ := range r.testCases {
		testCaseResult := r.validateTestCase(testCaseIndex)
		if !testCaseResult {
			testsResult = false
		}
	}

	if testsResult {
		r.runMainCase()
	}
}

func (r *Runner) Test(options ...runCaseOption) {
	r.testCases = append(r.testCases, runCaseWithOptions(options...))
}

func (r *Runner) Main(options ...runCaseOption) {
	r.mainCase = runCaseWithOptions(options...)
}

func (r *Runner) validateTestCase(testCaseIndex int) bool {
	fmt.Printf("Running test case %d\n", testCaseIndex+1)
	testCase := r.testCases[testCaseIndex]

	actualTestOutput := r.runFunc(testCase.InputLines())
	testResult := testCase.ValidateOutput(actualTestOutput)

	if testResult {
		fmt.Printf("\tTest result: succeeded\n")
	} else {
		fmt.Printf("\tTest result: failed\n")
	}

	fmt.Printf("\tActual output:   %s\n", actualTestOutput)
	fmt.Printf("\tExpected output: %s\n", testCase.output)

	return testResult
}

func (r *Runner) runMainCase() {
	fmt.Printf("Running main case\n")

	runOutput := r.runFunc(r.mainCase.InputLines())
	expectedResultStr := ""
	if r.mainCase.output != "" {
		runResult := r.mainCase.ValidateOutput(runOutput)
		if runResult {
			expectedResultStr = " (expected result)"
		} else {
			expectedResultStr = " (!!!non-expected result!!!)"
		}
	}
	fmt.Printf("Output: %s%s\n", runOutput, expectedResultStr)
}

func resolveTaskRootDirectory() string {
	_, currentFilePath, _, ok := runtime.Caller(2)
	if !ok {
		panic("failed to retrieve the task file path")
	}

	if !strings.HasSuffix(currentFilePath, "/main.go") {
		panic("NewRunner must be called from tasks/<task_name>/main.go")
	}

	taskRootDirectory := path.Dir(currentFilePath)
	return taskRootDirectory
}
