package main

import "stalary/study"

// append even list to odd list
func oddEvenList(head *study.ListNode) *study.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// odd list
	odd := head
	// even head list, use it append to odd list
	evenHead := head.Next
	// even list
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	// append even head list
	odd.Next = evenHead
	return head
}
