package main

import (
	"fmt"
	"strings"
)

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			high = mid - 1
		}

		if target > nums[mid] {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	numsLists := [][]int{
		{},
		{0, 1},
		{1, 2, 3},
		{-1, 0, 3, 5, 9, 12},
		{-100, -67, -55, -50, -49, -40, -33, -22, -10, -5},
	}

	targetList := []int{12, 1, 3, 9, -22}
	for i := 0; i < len(numsLists); i++ {
		nums := numsLists[i]
		target := targetList[i]
		index := binarySearch(nums, target)
		fmt.Printf("%d.\tArray to search: %v\n", i+1, strings.Replace(fmt.Sprint(nums), " ", ", ", -1))

		fmt.Printf("\tTarget: %d\n", target)
		if index != -1 {
			fmt.Printf("\t%d exists in the slice at index %d\n", target, index)
		} else {
			fmt.Printf("\t%d does not exist in the slice, so the return value is %d\n", target, index)
		}
		fmt.Println(strings.Repeat("-", 100))
	}
}
