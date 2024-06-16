package main

import (
	"fmt"
	"strings"
)

func reverseLinkedList(head *LinkedListNode) *LinkedListNode {
	var prev, next *LinkedListNode
	curr := head
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	head = prev
	return prev
}

func main() {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{3, 2, 1},
		{10},
		{1, 2},
	}
	for i := 0; i < len(input); i++ {
		inputLinkedList := LinkedList{}
		inputLinkedList.CreateLinkedList(input[i])
		DisplayLinkedList(inputLinkedList.head)
		fmt.Println("\nReversed Linked List: ")
		DisplayLinkedList(reverseLinkedList(inputLinkedList.head))
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
