package main

import (
	"fmt"
	"strings"
)

type pair struct {
	char  byte
	index int
}

func minRemoveParentheses(s string) string {
	stack := make([]pair, 0)
	sList := []byte(s)

	for i := 0; i < len(sList); i++ {
		val := sList[i]

		// if stack is not empty and top element of stack is an opening parenthesis
		// and the current element is a closing parenthesis
		if len(stack) > 0 && stack[len(stack)-1].char == '(' && val == ')' {
			// pop the opening parenthesis as it makes a valid pair
			// with the current closing parenthesis
			stack = stack[:len(stack)-1]
		} else if val == '(' || val == ')' {
			stack = append(stack, pair{char: val, index: i})
		}
	}

	// Remove the invalid parentheses
	for len(stack) > 0 {
		index := stack[len(stack)-1].index
		sList[index] = ' '
		stack = stack[:len(stack)-1]
	}

	var result strings.Builder
	for _, char := range sList {
		if char != ' ' {
			result.WriteByte(char)
		}
	}

	return result.String()
}

// Driver code
func main() {
	inputs := []string{"ar)ab(abc)abd(", "a)rt)lm(ikgh)", "aq)xy())qf(a(ba)q)", "(aw))kk())(w(aa)(bv(wt)r)", "(qi)(kl)((y(yt))(r(q(g)s)"}
	for i, input := range inputs {
		fmt.Printf("%d.\tInput: \"%s\"\n", i+1, input)
		fmt.Printf("\tValid parentheses, after minimum removal: \"%s\"\n", minRemoveParentheses(input))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
