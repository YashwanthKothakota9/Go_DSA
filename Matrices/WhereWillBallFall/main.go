package main

import (
	"fmt"
	"strings"
)

func FindExitColumn(matrix [][]int) []int {
	result := make([]int, len(matrix[0]))
	for i := range result {
		result[i] = -1
	}

	// Loop through each column in the grid
	for col := 0; col < len(matrix[0]); col++ {
		currColumn := col
		//Loop through each row in the grid
		for row := 0; row < len(matrix); row++ {
			nextColumn := currColumn + matrix[row][currColumn]
			fmt.Println("Next Column: ", nextColumn)

			// Check if the ball is out of the grid or hit a "V" shaped pattern
			if nextColumn < 0 || nextColumn > len(matrix[0])-1 || matrix[row][currColumn] != matrix[row][nextColumn] {
				break
			}

			//Check if ball reached the bottom
			if row == len(matrix)-1 {
				result[col] = nextColumn
			}
			currColumn = nextColumn
		}
	}

	return result
}

// Helper function to print a matrix
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, element := range row {
			fmt.Printf("%d\t", element)
		}
		fmt.Println()
	}
}

// Driver code
func main() {
	grids := [][][]int{
		{{1, 1, 1, -1}, {-1, -1, 1, 1}, {1, 1, -1, -1}, {-1, -1, 1, 1}},
	}

	for i, grid := range grids {
		fmt.Printf("Test Case #%d\n\n", i+1)
		fmt.Println("Input grid:")
		printMatrix(grid)
		fmt.Println()
		fmt.Println("Output:", strings.Replace(fmt.Sprint(FindExitColumn(grid)), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
