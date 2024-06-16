package main

func circularArrayLoop(nums []int) bool {
	size := len(nums)
	for i := 0; i < size; i++ {
		slow, fast := i, i
		forward := nums[i] > 0
		for {
			// Move slow pointer to one step.
			slow = nextStep(slow, nums[slow], size)
			// If cycle is not possible, break the loop and start from next element.
			if isNotCycle(nums, forward, slow) {
				break
			}

			// First move of fast pointer.
			fast = nextStep(fast, nums[fast], size)
			// If cycle is not possible, break the loop and start from next element.
			if isNotCycle(nums, forward, fast) {
				break
			}

			// Second move of fast pointer.
			fast = nextStep(fast, nums[fast], size)
			// If cycle is not possible, break the loop and start from next element.
			if isNotCycle(nums, forward, fast) {
				break
			}

			// At any point, if fast and slow pointers meet each other,
			// it indicate that loop has been found, return True.
			if slow == fast {
				return true
			}
		}
	}
	return false
}

// function to calculate next step
func nextStep(pointer int, value int, size int) int {
	result := (pointer + value) % size
	if result < 0 {
		result += size
	}
	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// A function to detect a cycle doesn't exist
func isNotCycle(nums []int, prevDirection bool, pointer int) bool {
	// Set current direction to True if current element is positive, set False otherwise.
	currDirection := nums[pointer] >= 0

	// If current direction and previous direction is different or moving a pointer takes back to the same value,
	// then the cycle is not possible, we return True, otherwise return False.
	if (prevDirection != currDirection) || (abs(nums[pointer])%len(nums) == 0) {
		return true
	} else {
		return false
	}
}
