package main

type ListNode struct {
    Val int
    Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var point *ListNode
	var root *ListNode
    for {
    	var temp ListNode
    	if l1!=nil && l2!=nil {
    		if l1.Val < l2.Val {
    			temp = *l1
    			l1 = l1.Next
    		} else {
    			temp = *l2
    			l2 = l2.Next
    		}
    	} else {
    		if l1!= nil {
    			temp = *l1
    			l1 = l1.Next
    		} else {
    			if l2 != nil{
    				temp = *l2
    				l2 = l2.Next
    			} else {
    				break
    			}
    		}
    	}
    	if root==nil {
    		root = &temp
    	} else {
    		point.Next = &temp
    	}
    	point = &temp
    }
    return root
}

func main() {
	a := &ListNode{Val:4,Next:nil}
	b := &ListNode{Val:2,Next:a}
	c := &ListNode{Val:1,Next:b}
	d := &ListNode{Val:4,Next:nil}
	e := &ListNode{Val:3,Next:d}
	f := &ListNode{Val:1,Next:e}
	ans := mergeTwoLists(c,f)
	for ;ans!=nil;{
		println(ans.Val)
		ans = ans.Next
	}
}