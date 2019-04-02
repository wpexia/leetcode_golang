type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
 

func isSymmetricSame(left *TreeNode, right *TreeNode) bool {
	if left == nil {
		if right == nil {
			return true
		} else {
			return false
		}
	} else if right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricSame(left.Left,right.Right) && isSymmetricSame(left.Right,right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left
	right := root.Right
	return isSymmetricSame(left,right)
}