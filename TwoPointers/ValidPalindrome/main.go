package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
	left := 0
	right := len(input) - 1
	for left < right {
		if input[left] != input[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	str := []string{"racecar", "a", "aba", "abcdefgfedcba", "ba"}
	for i, s := range str {
		fmt.Printf("Test case %d\n", i+1)
		fmt.Printf("%s\n", strings.Repeat("-", 10))
		fmt.Printf("The input string is %s\n", s)
		fmt.Printf("Is it a palindrome? %v\n", isPalindrome(s))
		fmt.Printf("%s\n", strings.Repeat("-", 10))
	}
}
