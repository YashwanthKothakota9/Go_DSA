// Template for Queue

package main

type Set struct {
	root *TreeNode[int]
	data int
}

// Node type struct
type Node struct {
	value *Set
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
func (q *Queue) Peek() (*TreeNode[int], int) {
	if q.IsEmpty() {
		return nil, -1
	}
	return q.head.value.root, q.head.value.data
}

// Enqueue push an element into the queue
func (q *Queue) Enqueue(root *TreeNode[int], data int) {
	value := &Set{root, data}
	temp := &Node{value, nil}
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
func (q *Queue) Dequeue() (*TreeNode[int], int) {
	if q.IsEmpty() {
		return nil, -1
	}
	value := q.head.value
	q.head = q.head.next
	q.size--
	return value.root, value.data
}
