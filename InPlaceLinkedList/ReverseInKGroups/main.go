package main

import (
	"fmt"
	"strings"
)

func reverseLinkedList(head *LinkedListNode, k int) (*LinkedListNode, *LinkedListNode) {
	var prev, curr, next *LinkedListNode = nil, head, nil
	for i := 0; i < k; i++ {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev, curr
}

func reverseKGroups(head *LinkedListNode, k int) *LinkedListNode {
	dummy := &LinkedListNode{data: 0}
	dummy.next = head
	ptr := dummy

	for ptr != nil {
		tracker := ptr
		for i := 0; i < k; i++ {
			if tracker == nil {
				break
			}
			tracker = tracker.next
		}
		if tracker == nil {
			break
		}
		prev, curr := reverseLinkedList(ptr.next, k)
		lastNodeOfReversedGroup := ptr.next
		lastNodeOfReversedGroup.next = curr
		ptr.next = prev
		ptr = lastNodeOfReversedGroup
	}
	return dummy.next
}

// Driver code
func main() {

	inputList := [][]int{{1, 2, 3, 4, 5, 6, 7, 8}, {3, 4, 5, 6, 2, 8, 7, 7}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5, 6, 7}, {1}}
	k := []int{3, 2, 1, 7, 1}

	for i := 0; i < len(inputList); i++ {

		var inputLinkedList LinkedList
		inputLinkedList.CreateLinkedList(inputList[i])

		fmt.Print(i+1, ".\tLinked list: ")
		DisplayLinkedList(inputLinkedList.head)
		fmt.Println()

		fmt.Print("\n\tReversed linked list: ")
		result := reverseKGroups(inputLinkedList.head, k[i])
		DisplayLinkedList(result)
		fmt.Println()

		fmt.Println(strings.Repeat("-", 100))
	}
}
