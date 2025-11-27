package runner

import (
	"fmt"
	"os"
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
	testPattern := filepath.Join(solutionPath, "test_*.txt")
	testFiles, err := filepath.Glob(testPattern)
	if err != nil {
		t.Fatalf("error finding test files: %v", err)
	}

	if len(testFiles) == 0 {
		t.Fatal("no test files found")
	}

	for _, testFile := range testFiles {
		fileName := filepath.Base(testFile)
		t.Run(fileName, func(t *testing.T) {
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

func readMainInputFile(solutionFolder string) ([]string, error) {
	inputFilePath := filepath.Join(solutionFolder, "main.in")

	if _, err := os.Stat(inputFilePath); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("input file does not exist: %s", inputFilePath)
		} else {
			return nil, fmt.Errorf("failed to check file stats: %v", err)
		}
	}

	inputFileRaw, err := os.ReadFile(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %v", err)
	}

	inputLines := strings.Split(string(inputFileRaw), "\n")

	return inputLines, nil
}

func readTestFile(testFilePath string) ([]string, string, error) {
	if _, err := os.Stat(testFilePath); err != nil {
		if os.IsNotExist(err) {
			return nil, "", fmt.Errorf("test file does not exist: %s", testFilePath)
		} else {
			return nil, "", fmt.Errorf("failed to check file stats: %v", err)
		}
	}

	testFileRaw, err := os.ReadFile(testFilePath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read file content: %v", err)
	}

	testFileLines := strings.Split(string(testFileRaw), "\n")
	testFilesLinesCount := len(testFileLines)

	separatorLineIndex := -1
	for i := testFilesLinesCount - 1; i >= 0; i-- {
		if isTestSeparatorLine(testFileLines[i]) {
			separatorLineIndex = i
			break
		}
	}

	if separatorLineIndex == -1 {
		return nil, "", fmt.Errorf("invalid test file format: missing separator line")
	}
	if separatorLineIndex < 1 {
		return nil, "", fmt.Errorf("invalid test file format: no input lines found")
	}
	if separatorLineIndex+1 >= testFilesLinesCount {
		return nil, "", fmt.Errorf("invalid test file format: missing expected output line")
	}

	testInputLines := testFileLines[:separatorLineIndex]
	outputLine := strings.TrimSpace(testFileLines[separatorLineIndex+1])

	return testInputLines, outputLine, nil
}

func isTestSeparatorLine(line string) bool {
	return len(line) > 0 && strings.HasPrefix(line, "=") && strings.HasSuffix(line, "=") && strings.Trim(line, "=") == ""
}
