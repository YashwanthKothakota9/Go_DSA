package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type RandomPickWithWeight struct {
	// List to store running sums of weights
	runningSums []int
	totalSum    int
}

func Constructor(weights []int) RandomPickWithWeight {
	// Variable to calculate running sum
	runningSum := 0
	runningSums := make([]int, len(weights))

	// Iterate through the given weights
	for i, w := range weights {
		// Add the current weight to the running sum
		runningSum += w
		// Append the running sum to the runningSums list
		runningSums[i] = runningSum
	}

	rand.Seed(time.Now().UnixNano())
	// Store the total sum of weights
	return RandomPickWithWeight{
		runningSums: runningSums,
		totalSum:    runningSum,
	}
}

// Method to pick an index based on the weights
func (r *RandomPickWithWeight) PickIndex() int {
	// Generate a random number between 1 and the total sum of the array
	target := rand.Intn(r.runningSums[len(r.runningSums)-1]) + 1
	// Initialize low and high variables for binary search
	low, high := 0, len(r.runningSums)

	// Perform binary search to find the first value higher than the target
	for low < high {
		mid := low + (high-low)/2
		if target > r.runningSums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	// Return the index (low) found
	return low
}

// Driver code
func main() {
	counter := 900
	weights := [][]int{
		{1, 2, 3, 4, 5},
		{1, 12, 23, 34, 45, 56, 67, 78, 89, 90},
		{10, 20, 30, 40, 50},
		{1, 10, 23, 32, 41, 56, 62, 75, 87, 90},
		{12, 20, 35, 42, 55},
		{10, 10, 10, 10, 10},
		{10, 10, 20, 20, 20, 30},
		{1, 2, 3},
		{10, 20, 30, 40},
		{5, 10, 15, 20, 25, 30},
	}
	dict := make(map[int]int)
	for index, value := range weights {
		obj := Constructor(value)
		fmt.Printf("%d.\tInput: %s, pickIndex() called %d times\n\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1), counter)
		for i, _ := range value {
			dict[i] = 0
		}
		for j := 0; j < counter; j++ {
			sol := obj.PickIndex()
			dict[sol] += 1
		}
		sum := 0
		for _, v := range value {
			sum += v
		}
		fmt.Printf("%s\n", strings.Repeat("-", 72))
		w := tabwriter.NewWriter(os.Stdout, 10, 0, 0, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "Indexes", "\t", "Weights", "\t", "Occurences", "\t", "Frequency", "\t", "Expected Frequency")
		fmt.Fprintln(w, strings.Repeat("-", 10), "\t", strings.Repeat("-", 10), "\t", strings.Repeat("-", 10), "\t", strings.Repeat("-", 10), "\t", strings.Repeat("-", 20))
		i := 0
		for _, v := range dict {
			if len(value) <= i {
				break
			}
			frequency := math.Round((float64(v)/float64(counter))*10000.0) / 100.0
			eFrequency := math.Round((float64(value[i])/float64(sum))*10000.0) / 100.0
			fmt.Fprintln(w, i, "\t", weights[index][i], "\t", v, "\t", frequency, "%", "\t", eFrequency, "%")
			i++
		}
		w.Flush()
		fmt.Printf("\n%s\n\n", strings.Repeat("-", 100))
	}
}
