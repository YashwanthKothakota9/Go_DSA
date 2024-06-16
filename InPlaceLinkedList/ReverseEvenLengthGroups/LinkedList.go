package main

import "fmt"

type LinkedList struct {
	head *LinkedListNode
}

// Inserts node at head
func (l *LinkedList) InsertNodeAtHead(node *LinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

// Creates linked list with given integer array
func (l *LinkedList) CreateLinkedList(input []int) {
	for i := len(input) - 1; i >= 0; i-- {
		newNode := InitLinkedListNode(input[i])
		l.InsertNodeAtHead(newNode)
	}
}

// Display LinkedList
func DisplayLinkedList(l *LinkedListNode) {
	temp := l
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(" -> ")
		} else {
			fmt.Print(" -> null")
		}
	}
}
