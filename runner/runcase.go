package runner

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

type runCase struct {
	input  string
	output string
}

func runCaseWithOptions(options ...runCaseOption) *runCase {
	runCase := &runCase{}
	for _, option := range options {
		option(runCase)
	}

	return runCase
}

func (t *runCase) InputLines() []string {
	return strings.Split(t.input, "\n")
}

func (t *runCase) ValidateOutput(actualOutput string) bool {
	actualOutput = strings.TrimSpace(actualOutput)
	expectedOutput := strings.TrimSpace(t.output)

	return actualOutput == expectedOutput
}

type runCaseOption func(*runCase)

func (r *Runner) InputFromText(text string) runCaseOption {
	return func(c *runCase) {
		c.input = text
	}
}

func (r *Runner) InputFromFile(file string) runCaseOption {
	return func(c *runCase) {
		var err error
		if c.input, err = readFileOutput(r.taskRootDirectory, file); err != nil {
			panic(fmt.Errorf("failed to read input file %s: %v", file, err))
		}
	}
}

func (r *Runner) OutputFromText(text string) runCaseOption {
	return func(c *runCase) {
		c.output = text
	}
}

func (r *Runner) OutputFromFile(file string) runCaseOption {
	return func(c *runCase) {
		var err error
		if c.output, err = readFileOutput(r.taskRootDirectory, file); err != nil {
			panic(fmt.Errorf("failed to read output file %s: %v", file, err))
		}
	}
}

func readFileOutput(taskRootDirectory, file string) (string, error) {
	fullFilePath := path.Join(taskRootDirectory, file)
	if _, err := os.Stat(fullFilePath); errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("file is not found: %s", fullFilePath)
	}

	fileBytes, err := os.ReadFile(fullFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return string(fileBytes), nil
}
