package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := head.Next
	last := head
	for temp != nil {
		if temp.Val != last.Val {
			last.Next = temp
			last = temp
		}
		temp = temp.Next
	}
	last.Next = nil
	return head
}
