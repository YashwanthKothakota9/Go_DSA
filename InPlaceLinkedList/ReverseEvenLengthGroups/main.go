package main

import (
	"fmt"
	"strings"
)

func reverseEvenLengthGroups(head *LinkedListNode) {
	prev := head
	groupLen := 2
	for prev.next != nil {
		node := prev
		numNodes := 0
		for i := 0; i < groupLen; i++ {
			if node.next == nil {
				break
			}
			numNodes += 1
			node = node.next
		}
		// If length of the current group is even
		if numNodes%2 == 0 {
			reverse := node.next
			curr := prev.next
			for i := 0; i < numNodes; i++ {
				currNext := curr.next
				curr.next = reverse
				reverse = curr
				curr = currNext
			}
			// updating the prev pointer after reversal of the even group
			prevNext := prev.next
			prev.next = node
			prev = prevNext
		} else {
			// If length of the current group is odd
			prev = node
		}
		// increment 1 by one and move to the next group
		groupLen += 1
	}
}

func main() {
	inputList := [][]int{
		{1, 2, 3, 4},
		{10, 11, 12, 13, 14},
		{15},
		{16, 17},
	}

	for index, value := range inputList {
		linkedList := new(LinkedList)
		linkedList.CreateLinkedList(value)
		fmt.Printf("%d.\tIf we reverse the even length groups of the linked list\n\t\t", index+1)
		DisplayLinkedList(linkedList.head)
		fmt.Printf("\n\tWe will get:\n\t\t")
		reverseEvenLengthGroups(linkedList.head)
		DisplayLinkedList(linkedList.head)
		fmt.Print(strings.Repeat("-", 100), "\n")
	}
}
