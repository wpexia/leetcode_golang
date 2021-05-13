package p0002

// ListNode is singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := new(ListNode)
	for current := dummy;l1 != nil || l2 != nil || carry != 0;current = current.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		current.Next =  &ListNode{Val: carry %10}
		carry /=10
	}
	return dummy.Next
}
