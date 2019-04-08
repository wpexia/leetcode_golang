package main
type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	newNode := &ListNode{}
	println(newNode)
	println()
	for head != nil {
		println(head)
		if head == newNode {
			return true
		}
		newHead := head.Next
		head.Next = newNode
		head = newHead
	}
	return false
}


func main(){
	Node1:= &ListNode{Val:2,Next:nil}
	Node2:= &ListNode{Val:0,Next:nil}
	Node3:= &ListNode{Val:-4,Next:nil}
	Node1.Next = Node2
	Node2.Next = Node3
	Node3.Next = Node1
	println(hasCycle(Node1))
}