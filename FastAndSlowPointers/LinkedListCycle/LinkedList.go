package main

import "fmt"

type LinkedList struct {
	head *LinkedListNode
}

// Insert node at head of a linked list
func (l *LinkedList) InsertNodeAtHead(node *LinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

// Creates linked list with given array of integers
func (l *LinkedList) CreateLinkedList(input []int) {
	for i := len(input) - 1; i >= 0; i-- {
		newNode := InitLinkedListNode(input[i])
		l.InsertNodeAtHead(newNode)
	}
}

// Length of linkedlist
func (l *LinkedList) GetLength() int {
	temp := l.head
	length := 0
	for temp != nil {
		length++
		temp = temp.next
	}
	return length
}

// returns node at specified position
func (l *LinkedList) GetNode(pos int) *LinkedListNode {
	if pos == -1 {
		return nil
	}
	p := 0
	ptr := l.head
	for p < pos {
		ptr = ptr.next
		p++
	}
	return ptr
}

// Display the elements of a linked list
func (l *LinkedList) DisplayLinkedList() {
	temp := l.head
	fmt.Print("[")
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print(']')
}

// Display elements of linked list with ->
func (l *LinkedList) DisplayLinkedListWithArrow() {
	temp := l.head
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
