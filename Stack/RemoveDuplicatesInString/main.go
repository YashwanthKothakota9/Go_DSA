package main

import (
	"fmt"
	"strings"
)

func removeDuplicates(str string) string {
	stack := make([]rune, 0)
	for _, c := range str {
		// If stack has at least one character and
		// stack's top character is same as the string's character
		if len(stack) > 0 && stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func main() {
	inputs := []string{
		"g",
		"ggaabcdeb",
		"abbddaccaaabcd",
		"aannkwwwkkkwna",
		"abbabccblkklu",
	}
	for i, input := range inputs {
		fmt.Printf("%d.\tRemove duplicates from string: \"%s\"\n", i+1, input)
		resultingString := removeDuplicates(input)
		fmt.Printf("\tString after removing duplicates: \"%s\"\n", resultingString)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
