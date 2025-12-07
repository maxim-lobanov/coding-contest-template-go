package main

import (
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	matrix := cast.ParseCharMatrix(input)

	dp := make([][]int, len(matrix))
	dp[0] = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix)-1; i++ {
		dp[i+1] = make([]int, len(matrix[i+1]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'S' {
				dp[i][j] = 1
			}

			if dp[i][j] == 0 {
				continue
			}

			if matrix[i+1][j] == '.' {
				dp[i+1][j] += dp[i][j]
			} else if matrix[i+1][j] == '^' {
				if j > 0 && matrix[i+1][j-1] == '.' {
					dp[i+1][j-1] += dp[i][j]
				}
				if j+1 < len(matrix[i]) && matrix[i+1][j+1] == '.' {
					dp[i+1][j+1] += dp[i][j]
				}
			}
		}
	}

	result := 0
	lastRow := dp[len(dp)-1]
	for i := 0; i < len(lastRow); i++ {
		result += lastRow[i]
	}

	return cast.ToString(result)
}
