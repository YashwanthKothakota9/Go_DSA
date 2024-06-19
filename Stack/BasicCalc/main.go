package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func calculator(expression string) int {
	signValue, number, result, popSignValue := 1, 0, 0, 0
	operationsStack := new(Stack)

	for _, c := range expression {
		if unicode.IsDigit(c) {
			temp2, _ := strconv.Atoi(string(c))
			number = number*10 + temp2
		}
		if c == '+' || c == '-' {
			result += number * signValue
			if c == '-' {
				signValue = -1
			} else {
				signValue = 1
			}
			number = 0
		} else if c == '(' {
			operationsStack.Push(result)
			operationsStack.Push(signValue)
			result = 0
			signValue = 1
		} else if c == ')' {
			result += signValue * number
			popSignValue = operationsStack.Top()
			operationsStack.Pop()
			result *= popSignValue
			secondValue := operationsStack.Top()
			operationsStack.Pop()
			result += secondValue
			number = 0
		}
	}
	return result + number*signValue
}

func main() {
	inputs := []string{"4 + (52 - 12) + 99",
		"(31 + 7) - (5 - 2)",
		"(12 - 9 + 4) + ( 7 - 5)",
		"8 - 5 + (19 - 11) + 6 + (10 + 3)",
		"56 - 44 - (27 - 17 - 1) + 7"}
	for i, input := range inputs {
		fmt.Printf("%d.\tGiven Expression: %s\n", i+1, input)
		fmt.Printf("\tThe result is: %d\n", calculator(input))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
