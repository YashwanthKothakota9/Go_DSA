package main

func getMiddleNode(head *LinkedListNode) *LinkedListNode {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
