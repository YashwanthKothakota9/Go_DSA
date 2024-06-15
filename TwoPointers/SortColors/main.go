package main

import (
	"fmt"
	"strings"
)

func sortColors(colors []int) []int {
	start := 0
	current := 0
	end := len(colors) - 1
	for current <= end {
		if colors[current] == 0 {
			if colors[start] != 0 {
				colors[start], colors[current] = colors[current], colors[start]
			}
			current++
			start++
		} else if colors[current] == 1 {
			current++
		} else {
			if colors[end] != 2 {
				colors[end], colors[current] = colors[current], colors[end]
			}
			end--
		}
	}
	return colors
}
func main() {
	input := []int{2, 1, 1, 0, 1, 0, 2}
	fmt.Printf("colors: %v\n", arrayToString(input))
	sortedColors := sortColors(input)
	fmt.Printf("Sorted Colros: %v\n", arrayToString(sortedColors))
}

func arrayToString(arr []int) string {
	strArr := make([]string, len(arr))
	for i, num := range arr {
		strArr[i] = fmt.Sprint(num)
	}
	return "[" + strings.Join(strArr, ", ") + "]"
}
