package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func readMainInputFile(solutionFolder string) ([]string, error) {
	inputFilePath := filepath.Join(solutionFolder, "input.txt")

	if _, err := os.Stat(inputFilePath); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("error reading input file: file does not exist: %s", inputFilePath)
		} else {
			return nil, fmt.Errorf("error checking file stats: %w", err)
		}
	}

	inputFileRaw, err := os.ReadFile(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file content: %w", err)
	}

	inputLines := strings.Split(string(inputFileRaw), "\n")

	return inputLines, nil
}

func findAllTestFiles(solutionFolder string) ([]string, error) {
	testPattern := filepath.Join(solutionFolder, "sample_*.txt")
	testFiles, err := filepath.Glob(testPattern)
	if err != nil {
		return nil, fmt.Errorf("error listing files by glob: %w", err)
	}

	return testFiles, nil
}

func readTestFile(testFilePath string) ([]string, string, error) {
	if _, err := os.Stat(testFilePath); err != nil {
		if os.IsNotExist(err) {
			return nil, "", fmt.Errorf("error reading test file: file does not exist: %s", testFilePath)
		} else {
			return nil, "", fmt.Errorf("error checking file stats: %w", err)
		}
	}

	testFileRaw, err := os.ReadFile(testFilePath)
	if err != nil {
		return nil, "", fmt.Errorf("error reading file content: %w", err)
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
		return nil, "", fmt.Errorf("error parsing test file: missing separator line")
	}
	if separatorLineIndex < 1 {
		return nil, "", fmt.Errorf("error parsing test file: no input lines found")
	}
	if separatorLineIndex+1 >= testFilesLinesCount {
		return nil, "", fmt.Errorf("error parsing test file: missing expected output line")
	}

	testInputLines := testFileLines[:separatorLineIndex]
	outputLine := strings.TrimSpace(testFileLines[separatorLineIndex+1])

	return testInputLines, outputLine, nil
}

func isTestSeparatorLine(line string) bool {
	return line == "---OUTPUT---"
}
