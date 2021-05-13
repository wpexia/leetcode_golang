package main
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type DeepTreeNode struct {
	Deep int
	Node *TreeNode
	Next *DeepTreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	head := &DeepTreeNode{Deep:1,Node:root,Next:nil}
	bottom := head
	for head!= nil {
		if (*(*head).Node).Left != nil {
			bottom.Next = &DeepTreeNode{Deep:(*head).Deep+1,Node:(*(*head).Node).Left,Next:nil}
			bottom = bottom.Next
		}
		if (*(*head).Node).Right != nil {
			bottom.Next = &DeepTreeNode{Deep:(*head).Deep+1,Node:(*(*head).Node).Right,Next:nil}
			bottom = bottom.Next
		}
		if (*(*head).Node).Left == nil && (*(*head).Node).Right == nil {
			return (*head).Deep
		} else {
			head = (*head).Next
		}
	}
	return 0
}

func ArrayToTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val:nums[0],Left:nil,Right:nil}
	nodeArray := make([]*TreeNode, len(nums))
	nodeArray[0] = root
	for i:=1;i<len(nums);i++{
		if nums[i] <0 {
			continue
		}
		node := &TreeNode{Val:nums[i],Left:nil,Right:nil}
		nodeArray[i] = node
		if i % 2 !=0 {
			(*(nodeArray[(i-1)/2])).Left = node
		} else {
			(*(nodeArray[(i-1)/2])).Right = node
		}
	}
	return root
}


func main() {
	root := ArrayToTree([]int{1,2})
	println(minDepth(root))
}