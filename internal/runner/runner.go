package runner

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SolutionFunc func(input []string) string

func Execute(t *testing.T, solutionFn SolutionFunc) {
	solutionPath, solutionRelativePath, err := getSolutionPath()
	if err != nil {
		t.Fatalf("error getting solution path: %v", err)
	}

	t.Logf("Solution path: %s", solutionRelativePath)

	executeTests(t, solutionFn, solutionPath)

	if !t.Failed() {
		executeMain(t, solutionFn, solutionPath)
	}

}

func executeMain(t *testing.T, solutionFn SolutionFunc, solutionPath string) {
	t.Run("main", func(t *testing.T) {
		input, err := readMainInputFile(solutionPath)
		if err != nil {
			t.Fatalf("error reading main input file: %v", err)
		}

		defer enableRecovery(t)
		actualOutput := solutionFn(input)

		t.Log("Result for main input:")
		t.Log(actualOutput)
		assert.NotEmpty(t, actualOutput)
	})
}

func executeTests(t *testing.T, solutionFn SolutionFunc, solutionPath string) {
	testFiles, err := findAllTestFiles(solutionPath)
	if err != nil {
		t.Fatalf("error finding test files: %v", err)
	}

	if len(testFiles) == 0 {
		t.Fatal("no test files found")
	}

	for _, testFile := range testFiles {
		testName := strings.TrimSuffix(filepath.Base(testFile), filepath.Ext(testFile))
		t.Run(testName, func(t *testing.T) {
			testInput, expectedOutput, err := readTestFile(testFile)
			if err != nil {
				t.Fatalf("error reading test file: %v", err)
			}

			defer enableRecovery(t)
			actualOutput := solutionFn(testInput)

			if actualOutput != expectedOutput {
				if !assert.Equal(t, expectedOutput, actualOutput) {
					t.Fail()
				}
			}
		})
	}
}

func enableRecovery(t *testing.T) {
	if r := recover(); r != nil {
		t.Fatalf("solution function panicked: %v", r)
	}
}
