package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func readMainInputFile(solutionFolder string) ([]string, error) {
	inputFilePath := filepath.Join(solutionFolder, "main.in")

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
	testInputPattern := filepath.Join(solutionFolder, "sample_*.in")
	testInputFiles, err := filepath.Glob(testInputPattern)
	if err != nil {
		return nil, fmt.Errorf("error listing files by glob: %w", err)
	}

	foundTests := make([]string, 0, len(testInputFiles))
	for _, testFile := range testInputFiles {
		baseTestFile := strings.TrimSuffix(testFile, ".in")
		expectedOutputFile := baseTestFile + ".out"
		if _, err := os.Stat(expectedOutputFile); err == nil {
			foundTests = append(foundTests, baseTestFile)
		}
	}

	return foundTests, nil
}

func readTestFile(baseTestFile string) ([]string, string, error) {
	testInputFile := fmt.Sprintf("%s.in", baseTestFile)
	testOutputFile := fmt.Sprintf("%s.out", baseTestFile)

	testInputRaw, err := os.ReadFile(testInputFile)
	if err != nil {
		return nil, "", fmt.Errorf("error reading input file: %w", err)
	}

	testOutputRaw, err := os.ReadFile(testOutputFile)
	if err != nil {
		return nil, "", fmt.Errorf("error reading output file: %w", err)
	}

	testInputLines := strings.Split(string(testInputRaw), "\n")
	testOutputLine := strings.TrimSpace(string(testOutputRaw))

	return testInputLines, testOutputLine, nil
}
