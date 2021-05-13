type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalancedAndDeepth(root *TreeNode) (bool,int) {
	if root == nil{
		return true,0
	}
	leftFlag,leftDeepth := isBalancedAndDeepth(root.Left)
	rightFlag,rightDeepth  := isBalancedAndDeepth(root.Right)
	if !(leftFlag && rightFlag) {
		return false,0
	}
	if leftDeepth > rightDeepth {
		if leftDeepth - rightDeepth >1 {
			return false,0
		} else {
			return true, leftDeepth + 1
		}
	} else {
		if rightDeepth - leftDeepth >1 {
			return false,0
		} else {
			return true, rightDeepth +1
		}
	}
}

func isBalanced(root *TreeNode) bool {
	ans,_ := isBalancedAndDeepth(root)
	return ans
}