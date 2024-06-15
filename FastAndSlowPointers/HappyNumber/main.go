package main

import (
	"fmt"
	"strings"
)

// calculates power of given digit
func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

func sumOfSquaredDigits(number int) int {
	totalSum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		totalSum += pow(digit, 2)
	}
	return totalSum
}

func isHappyNumber(number int) bool {
	slow := number
	fast := sumOfSquaredDigits(number)
	for fast != 1 && slow != fast {
		slow = sumOfSquaredDigits(slow)
		fast = sumOfSquaredDigits(sumOfSquaredDigits(fast))
	}
	return fast == 1
}

func main() {
	array := []int{5, 1, 19, 25, 2}
	for i, v := range array {
		fmt.Printf("\n%d. Input number: %d", (i + 1), v)
		fmt.Printf("\nIs it a happy number? %t", isHappyNumber(v))
		fmt.Printf("\n%s", strings.Repeat("-", 20))
	}
}
