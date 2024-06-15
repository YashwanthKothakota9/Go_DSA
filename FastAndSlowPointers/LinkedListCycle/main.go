package main

func detectCycle(head *LinkedListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}

	return false
}
