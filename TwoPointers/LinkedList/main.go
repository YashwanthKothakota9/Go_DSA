package main

import (
	"fmt"
	"strings"
)

func removeNthNodeFromEndOfLinkedList(head *LinkedListNode, n int) *LinkedListNode {
	right := head
	left := head

	for i := 0; i < n; i++ {
		right = right.next
	}

	//removal of head node
	if right == nil {
		return head.next
	}

	for right.next != nil {
		right = right.next
		left = left.next
	}

	//now left points at (n-1)th node
	left.next = left.next.next

	return head
}

func main() {
	inputs := [][]int{
		{23, 89, 10, 5, 67, 70, 28},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	n := []int{4, 1}
	for i := 0; i < len(inputs); i++ {
		inputLinkedList := LinkedList{}
		inputLinkedList.CreateLinkedList(inputs[i])
		fmt.Printf("\n%d. LinkedList:\t", (i + 1))
		inputLinkedList.DisplayLinkedListWithArrow()
		fmt.Printf("\nn:%d", n[i])
		fmt.Printf("\nUpdated LinkedList:\t")
		removeNthNodeFromEndOfLinkedList(inputLinkedList.head, n[i])
		inputLinkedList.DisplayLinkedListWithArrow()
		fmt.Printf("\n%s\n", strings.Repeat("-", 50))
	}
}
