package main

import (
	"fmt"
	"strings"
)

func reorderList(head *LinkedListNode) {
	if head == nil {
		return
	}
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow, fast = slow.next, fast.next.next
	}

	prev := new(LinkedListNode)
	prev = nil
	curr := slow
	temp := new(LinkedListNode)
	temp = nil
	for curr != nil {
		temp = curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	// merge two sorted linked lists
	// merge 1->2->3 and 6->5->4 into 1->6->2->5->3->4
	first, second := head, prev
	for second.next != nil {
		temp = first.next
		first.next = second
		first = temp

		temp = second.next
		second.next = first
		second = temp
	}
}

// Driver code
func main() {
	inputLists := [][]int{
		{1, 1, 2, 2, 3, -1, 10, 12},
		{10, 20, -22, 21, -12},
		{1, 1, 1},
		{-2, -5, -6, 0, -1, -4},
		{3, 1, 5, 7, -4, -2, -1, -6},
	}

	for i, inputList := range inputLists {
		obj := new(LinkedList)
		obj.CreateLinkedList(inputList)
		fmt.Printf("%d.\tOriginal list: ", i+1)
		DisplayLinkedList(obj.head)
		reorderList(obj.head)
		fmt.Print("\n\tAfter folding: ")
		DisplayLinkedList(obj.head)
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
