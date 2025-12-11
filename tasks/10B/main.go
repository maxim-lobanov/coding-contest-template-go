package main

import (
	"math"
	"strings"

	"github.com/maxim-lobanov/coding-contest-template-go/internal/algo"
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	totalResult := 0
	for _, line := range input {
		result := solveSingleCase(line)
		totalResult += result
	}
	return cast.ToString(totalResult)
}

func solveSingleCase(input string) int {
	requiredPattern, availableButtons := parseInputLine(input)
	numVars := len(availableButtons)
	numEqs := len(requiredPattern)

	// Build coefficient matrix
	matrix := make([][]float64, numEqs)
	for i := 0; i < numEqs; i++ {
		matrix[i] = make([]float64, numVars)
		for j := 0; j < numVars; j++ {
			for _, idx := range availableButtons[j] {
				if idx == i {
					matrix[i][j] = 1
				}
			}
		}
	}

	b := make([]float64, numEqs)
	for i := 0; i < numEqs; i++ {
		b[i] = float64(requiredPattern[i])
	}

	solution := solveSystem(matrix, b)

	total := 0
	for _, v := range solution {
		total += v
	}

	return total
}

func parseInputLine(input string) ([]int, [][]int) {
	parts := strings.Split(input, "]")
	parts = strings.Split(parts[1], "{")
	buttonsRaw := strings.Split(strings.TrimSpace(parts[0]), " ")

	requiredPatternRaw := strings.TrimSuffix(parts[1], "}")
	requiredPatternParts := strings.Split(requiredPatternRaw, ",")
	requiredPattern := algo.Map(requiredPatternParts, func(item string) int { return cast.ParseInt(item) })

	availableButtons := [][]int{}
	for _, buttonRawPart := range buttonsRaw {
		buttonRawPart = strings.TrimSpace(buttonRawPart)
		buttonRawPart = strings.TrimPrefix(buttonRawPart, "(")
		buttonRawPart = strings.TrimSuffix(buttonRawPart, ")")
		optionRawList := strings.Split(buttonRawPart, ",")
		optionsList := algo.Map(optionRawList, func(item string) int { return cast.ParseInt(item) })
		availableButtons = append(availableButtons, optionsList)
	}

	return requiredPattern, availableButtons
}

func solveSystem(A [][]float64, b []float64) []int {
	// For small systems, try a more direct approach with optimization
	m := len(A[0])

	// Try branch and bound / search for minimal solution
	if m <= 15 { // Small enough for smart enumeration
		return solveBranchAndBound(A, b)
	}

	return solveGaussian(A, b)
}

func solveGaussian(A [][]float64, b []float64) []int {
	n := len(A)
	m := len(A[0])

	// Augmented matrix
	aug := make([][]float64, n)
	for i := 0; i < n; i++ {
		aug[i] = make([]float64, m+1)
		for j := 0; j < m; j++ {
			aug[i][j] = A[i][j]
		}
		aug[i][m] = b[i]
	}

	// RREF
	lead := 0
	for r := 0; r < n; r++ {
		if lead >= m {
			break
		}
		i := r
		for math.Abs(aug[i][lead]) < 1e-10 {
			i++
			if i == n {
				i = r
				lead++
				if lead == m {
					break
				}
			}
		}
		if lead >= m {
			break
		}

		aug[i], aug[r] = aug[r], aug[i]
		lv := aug[r][lead]
		for j := 0; j <= m; j++ {
			aug[r][j] /= lv
		}
		for i := 0; i < n; i++ {
			if i != r {
				lv = aug[i][lead]
				for j := 0; j <= m; j++ {
					aug[i][j] -= lv * aug[r][j]
				}
			}
		}
		lead++
	}

	// Extract solution
	solution := make([]float64, m)

	// Identify pivot columns
	pivotCol := make([]int, n)
	basicVar := make(map[int]int) // maps variable index to its pivot row
	for i := 0; i < n; i++ {
		pivotCol[i] = -1
	}

	for row := 0; row < n; row++ {
		// Find the leading 1 in this row (pivot)
		for col := 0; col < m; col++ {
			if math.Abs(aug[row][col]-1.0) < 1e-9 {
				// Check if this is truly a pivot (only non-zero in this column)
				isPivot := true
				for otherRow := 0; otherRow < n; otherRow++ {
					if otherRow != row && math.Abs(aug[otherRow][col]) > 1e-9 {
						isPivot = false
						break
					}
				}
				if isPivot {
					pivotCol[row] = col
					basicVar[col] = row
					break
				}
			}
		}
	}

	// Identify free variables (non-basic variables)
	freeVars := []int{}
	for col := 0; col < m; col++ {
		if _, isBasic := basicVar[col]; !isBasic {
			freeVars = append(freeVars, col)
		}
	}

	// If there are free variables, we need to find values that keep all basic vars non-negative
	// and minimize the total sum
	if len(freeVars) > 0 {
		// For each free variable, compute the minimum value needed to keep all basic vars >= 0
		for _, freeVar := range freeVars {
			minVal := 0.0
			for basicVarIdx, row := range basicVar {
				coeff := aug[row][freeVar]
				rhs := aug[row][m]
				// Basic var = rhs - coeff * freeVar
				// We need: rhs - coeff * freeVar >= 0
				// If coeff > 0: freeVar <= rhs/coeff (no lower bound from this)
				// If coeff < 0: freeVar >= rhs/coeff (provides lower bound)
				if coeff < -1e-9 {
					requiredMin := rhs / coeff
					if requiredMin > minVal {
						minVal = requiredMin
					}
				}
				_ = basicVarIdx // use it
			}
			solution[freeVar] = math.Ceil(minVal) // Round up to ensure non-negativity
		}
	}

	// Now compute basic variables
	for basicVarIdx, row := range basicVar {
		val := aug[row][m]
		for _, freeVar := range freeVars {
			val -= aug[row][freeVar] * solution[freeVar]
		}
		solution[basicVarIdx] = math.Round(val)
	}

	// Convert to int
	intSolution := make([]int, m)
	for i := 0; i < m; i++ {
		intSolution[i] = int(solution[i])
	}

	return intSolution
}

func solveBranchAndBound(A [][]float64, b []float64) []int {
	numEqs := len(A)
	numVars := len(A[0])

	// First do Gaussian elimination to identify basic and free variables
	aug := make([][]float64, numEqs)
	for i := 0; i < numEqs; i++ {
		aug[i] = make([]float64, numVars+1)
		for j := 0; j < numVars; j++ {
			aug[i][j] = A[i][j]
		}
		aug[i][numVars] = b[i]
	}

	// RREF
	lead := 0
	for r := 0; r < numEqs; r++ {
		if lead >= numVars {
			break
		}
		i := r
		for math.Abs(aug[i][lead]) < 1e-10 {
			i++
			if i == numEqs {
				i = r
				lead++
				if lead == numVars {
					break
				}
			}
		}
		if lead == numVars {
			break
		}
		aug[i], aug[r] = aug[r], aug[i]

		if math.Abs(aug[r][lead]) > 1e-10 {
			div := aug[r][lead]
			for j := 0; j <= numVars; j++ {
				aug[r][j] /= div
			}
		}

		for i := 0; i < numEqs; i++ {
			if i != r {
				mult := aug[i][lead]
				for j := 0; j <= numVars; j++ {
					aug[i][j] -= mult * aug[r][j]
				}
			}
		}
		lead++
	}

	// Identify pivot columns (basic variables)
	basicVar := make(map[int]int) // maps variable index to its pivot row
	for row := 0; row < numEqs; row++ {
		for col := 0; col < numVars; col++ {
			if math.Abs(aug[row][col]-1.0) < 1e-9 {
				isPivot := true
				for otherRow := 0; otherRow < numEqs; otherRow++ {
					if otherRow != row && math.Abs(aug[otherRow][col]) > 1e-9 {
						isPivot = false
						break
					}
				}
				if isPivot {
					basicVar[col] = row
					break
				}
			}
		}
	}

	// Identify free variables
	freeVars := []int{}
	for col := 0; col < numVars; col++ {
		if _, isBasic := basicVar[col]; !isBasic {
			freeVars = append(freeVars, col)
		}
	}

	// Now try different combinations of free variable values
	bestSolution := make([]int, numVars)
	bestSum := 1000000

	// Compute upper bound for free variables based on the system
	maxValue := 0
	for i := 0; i < numEqs; i++ {
		if int(math.Ceil(aug[i][numVars])) > maxValue {
			maxValue = int(math.Ceil(aug[i][numVars]))
		}
	}
	if maxValue < 30 {
		maxValue = 30 //  Minimum reasonable search space
	}

	// Add buffer, but limit total search space
	searchBudget := 100000 // Maximum combinations to try
	if len(freeVars) > 0 {
		maxPerVar := int(math.Pow(float64(searchBudget), 1.0/float64(len(freeVars))))
		if maxPerVar > maxValue*3 {
			maxValue = maxValue * 3
		} else {
			maxValue = maxPerVar
		}
	}

	var tryFreeVars func(freeIdx int, freeAssignment []int)
	tryFreeVars = func(freeIdx int, freeAssignment []int) {
		if freeIdx == len(freeVars) {
			// Compute basic variables
			solution := make([]float64, numVars)
			for i, val := range freeAssignment {
				solution[freeVars[i]] = float64(val)
			}

			valid := true
			currentSum := 0
			for basicVarIdx, row := range basicVar {
				val := aug[row][numVars]
				for _, freeVar := range freeVars {
					val -= aug[row][freeVar] * solution[freeVar]
				}
				if val < -0.5 {
					valid = false
					break
				}
				roundedVal := math.Round(val)
				if math.Abs(val-roundedVal) > 1e-6 {
					valid = false
					break
				}
				solution[basicVarIdx] = roundedVal
				currentSum += int(roundedVal)
			}

			if valid {
				// Add free variable values to sum
				for _, v := range freeAssignment {
					currentSum += v
				}

				if currentSum < bestSum {
					bestSum = currentSum
					for i := range solution {
						bestSolution[i] = int(solution[i])
					}
				}
			}
			return
		}

		// Try values with pruning
		for val := 0; val <= maxValue; val++ {
			// Prune: if we already have too large a sum from free vars alone
			currentFreeSum := val
			for i := 0; i < freeIdx; i++ {
				currentFreeSum += freeAssignment[i]
			}
			if currentFreeSum >= bestSum {
				break // Prune this branch
			}

			freeAssignment[freeIdx] = val
			tryFreeVars(freeIdx+1, freeAssignment)
		}
	}

	if len(freeVars) > 0 {
		freeAssignment := make([]int, len(freeVars))
		tryFreeVars(0, freeAssignment)
	} else {
		// No free variables, just extract the solution
		for basicVarIdx, row := range basicVar {
			bestSolution[basicVarIdx] = int(math.Round(aug[row][numVars]))
		}
	}

	return bestSolution
}
