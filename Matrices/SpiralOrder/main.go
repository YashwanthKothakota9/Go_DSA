package main

import (
	"fmt"
	"strings"
)

func SpiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])

	//Pointers to traverse matrix
	row := 0
	col := -1

	//Initial direction
	direction := 1

	var result []int

	for rows > 0 && cols > 0 {
		//Left -> Right if direction==1
		//Right -> Left if direction==-1
		for i := 0; i < cols; i++ {
			col += direction
			result = append(result, matrix[row][col])
		}
		rows--

		//Top -> Bottom if direction==1
		//Bottom -> Top if direction==-1
		for i := 0; i < rows; i++ {
			row += direction
			result = append(result, matrix[row][col])
		}
		cols--

		//Flip direction
		direction *= -1
	}

	return result
}

// Driver code
func main() {
	inputs := [][][]int{
		{{1}},
		{{6}, {2}},
		{{2, 14, 8}, {12, 7, 14}},
		{{3, 1, 1}, {15, 12, 13}, {4, 14, 12}, {10, 5, 11}},
		{{10, 1, 14, 11, 14}, {13, 4, 8, 2, 13}, {10, 19, 1, 6, 8}, {20, 10, 8, 2, 12}, {15, 6, 8, 8, 18}},
	}

	for i, matrix := range inputs {
		fmt.Printf("%d.\tMatrix:\n", i+1)
		printMatrix(matrix)

		fmt.Printf("\n\tSpiral order: %s\n", strings.Replace(fmt.Sprint(SpiralOrder(matrix)), " ", ", ", -1))

		fmt.Println(strings.Repeat("-", 100))
	}
}
func printMatrix(matrix [][]int) {
	fmt.Print("\t\t[")
	rowNum := 0
	for i := 0; i < len(matrix); i++ {
		if rowNum == 0 {
			fmt.Print("[")
		} else {
			fmt.Print("\t\t [")
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
