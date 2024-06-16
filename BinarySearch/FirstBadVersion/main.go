package main

import (
	"fmt"
	"strings"
)

var v int

func isBadVersion(s int) bool {
	return s >= v
}

func firstBadVersion(n int) (int, int) {
	first := 1
	last := n
	calls := 0

	for first <= last {
		mid := first + (last-first)/2
		if isBadVersion(mid) {
			last = mid - 1
		} else {
			first = mid + 1
		}
		calls += 1
	}
	return first, calls
}

// Driver code
func main() {
	testCaseVersions := []int{38, 13, 29, 40, 23, 1000}
	firstBadVersions := []int{28, 10, 10, 28, 19, 999}

	for index, value := range testCaseVersions {
		v = firstBadVersions[index]
		if index > 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%d.\tNumber of versions: %d\n", index+1, value)
		result, result1 := firstBadVersion(value)
		fmt.Printf("\n\tFirst bad version: %d. Found in %d API calls.\n", result, result1)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
