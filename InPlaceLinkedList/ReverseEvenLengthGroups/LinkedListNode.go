package main

type LinkedListNode struct {
	data int
	next *LinkedListNode
}

func NewLinkedListNode(data int, next *LinkedListNode) *LinkedListNode {
	node := new(LinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *LinkedListNode {
	node := new(LinkedListNode)
	node.data = data
	node.next = nil
	return node
}
