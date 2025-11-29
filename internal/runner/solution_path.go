package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getSolutionPath() (string, string, error) {
	_, callerFile, _, ok := runtime.Caller(2)
	if !ok {
		return "", "", fmt.Errorf("error getting caller information")
	}

	solutionFullPath := filepath.Dir(callerFile)
	solutionRelativePath, err := getSolutionRelativePath(solutionFullPath)
	if err != nil {
		return "", "", fmt.Errorf("error getting solution relative path: %w", err)
	}

	return solutionFullPath, solutionRelativePath, nil
}

func getSolutionRelativePath(solutionPath string) (string, error) {
	currentDir := solutionPath

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("error finding repo root directory from path: %s", solutionPath)
		}
		currentDir = parentDir
	}

	solutionRelativePath, err := filepath.Rel(currentDir, solutionPath)
	if err != nil {
		return "", fmt.Errorf("error calculating relative path: %w", err)
	}

	return solutionRelativePath, nil
}
