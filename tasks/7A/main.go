package main

import (
	"github.com/maxim-lobanov/coding-contest-template-go/internal/cast"
)

func solution(input []string) string {
	matrix := cast.ParseCharMatrix(input)

	result := 0
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 'S' && matrix[i][j] != '|' {
				continue
			}

			if matrix[i+1][j] == '.' {
				matrix[i+1][j] = '|'
			} else if matrix[i+1][j] == '^' {
				result++
				if j > 0 && matrix[i+1][j-1] == '.' {
					matrix[i+1][j-1] = '|'
				}
				if j+1 < len(matrix[i+1]) && matrix[i+1][j+1] == '.' {
					matrix[i+1][j+1] = '|'
				}
			}
		}
	}

	/*
		for i := 0; i < len(matrix); i++ {
			fmt.Println(string(matrix[i]))
		}
	*/

	return cast.ToString(result)
}
