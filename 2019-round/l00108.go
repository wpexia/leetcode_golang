type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func buildTreeNode(nums *[]int,left int, right int) * TreeNode {
	if left > right {
		return nil
	}
	middle := (left+right)/2
	ans := &TreeNode{Val: (*nums)[middle],Left:nil,Right:nil}
	ans.Left = buildTreeNode(nums,left,middle-1)
	ans.Right = buildTreeNode(nums,middle+1,right)
	return ans
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildTreeNode(&nums,0,len(nums)-1)
}