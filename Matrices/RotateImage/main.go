package main

import (
	"fmt"
	"strings"
)

func rotateImage(matrix [][]int) [][]int {
	n := len(matrix)

	for row := 0; row < n/2; row++ {
		for col := row; col < n-row-1; col++ {
			// Swap the top-left and top-right cells in the current group
			matrix[row][col], matrix[col][n-1-row] = matrix[col][n-1-row], matrix[row][col]

			//Swap the top-left and bottom-right cells in the current group
			matrix[row][col], matrix[n-1-row][n-1-col] = matrix[n-1-row][n-1-col], matrix[row][col]

			// Swap the top-left and bottom-left cells in the current group
			matrix[row][col], matrix[n-1-col][row] = matrix[n-1-col][row], matrix[row][col]
		}
	}

	return matrix
}

func main() {
	inputs := [][][]int{
		{{1}},
		{{6, 9}, {2, 7}},
		{{2, 14, 8}, {12, 7, 14}, {3, 3, 7}},
		{{3, 1, 1, 7}, {15, 12, 13, 13}, {4, 14, 12, 4}, {10, 5, 11, 12}},
		{{10, 1, 14, 11, 14}, {13, 4, 8, 2, 13}, {10, 19, 1, 6, 8}, {20, 10, 8, 2, 12}, {15, 6, 8, 8, 18}},
	}
	for i, input := range inputs {

		fmt.Printf("%d.\tMatrix:\n", i+1)
		printMatrix(input)

		fmt.Println("\n\tRotated matrix:")
		printMatrix(rotateImage(input))

		fmt.Println(strings.Repeat("-", 100))
	}
}

func printMatrix(matrix [][]int) {
	fmt.Print("\t\t\t[")
	rowNum := 0
	for i := 0; i < len(matrix); i++ {
		if rowNum == 0 {
			fmt.Print("[")
		} else {
			fmt.Print("\t\t\t [")
		}
		rowNum++

		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
			if j != len(matrix[i])-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("]")
		if i != len(matrix)-1 {
			fmt.Println(",")
		}
	}
	fmt.Println("]")
}
