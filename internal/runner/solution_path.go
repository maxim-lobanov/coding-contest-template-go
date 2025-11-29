package runner

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func getSolutionPath() (string, error) {
	_, callerFile, _, ok := runtime.Caller(2)
	if !ok {
		return "", fmt.Errorf("error getting caller information")
	}

	solutionFullPath := filepath.Dir(callerFile)

	return solutionFullPath, nil
}
