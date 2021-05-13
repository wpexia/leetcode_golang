type ListNode struct {
	Val int
	Next *ListNode
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	for curA!=nil &&curB!=nil {
		curA = curA.Next
		curB = curB.Next
	}
	for curA != nil {
		curA = curA.Next
		headA = headA.Next
	}
	for curB != nil {
		curB = curB.Next
		headB = headB.Next
	}
	for headA!=headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}