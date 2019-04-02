package algorithm

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 负数代表空节点
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