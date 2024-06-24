package main

import (
	"fmt"
	"strconv"
)

// Queue is used by binary tree

// Node type struct
type Node struct {
	value *TreeNode[int]
	next  *Node
}

// Queue type struct
type Queue struct {
	head *Node
	tail *Node
	size int
}

// Size returns the size of the queue
func (q *Queue) Size() int {
	return q.size
}

// IsEmpty checks whether the queue is empty or not
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Peek returns the top element of the queue
func (q *Queue) Peek() *TreeNode[int] {
	if q.IsEmpty() {
		return nil
	}
	return q.head.value
}

// Enqueue push an element into the queue
func (q *Queue) Enqueue(data *TreeNode[int]) {
	temp := &Node{data, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

// Dequeue remove an element from the queue
func (q *Queue) Dequeue() *TreeNode[int] {
	if q.IsEmpty() {
		return nil
	}
	value := q.head.value
	q.head = q.head.next
	q.size--
	return value
}

// String function converts queue elements to string
func (q *Queue) String() string {
	output := "["
	if q.head != nil {
		temp := q.head
		for temp != nil {
			fmt.Println(output)
			output += strconv.Itoa(temp.value.Data) + ", "
			temp = temp.next
		}
		output = output[0 : len(output)-2]
	}
	output += "]"
	return output
}
