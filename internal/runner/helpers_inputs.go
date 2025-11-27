package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func readMainInputFile(solutionFolder string) ([]string, error) {
	inputFilePath := filepath.Join(solutionFolder, "main.txt")

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
