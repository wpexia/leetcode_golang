package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func orderBottom(point *TreeNode,length int, ans *[][]int) {
	if point == nil {
		return
	}
	if len(*ans) <=length{
		newLevel:= make([]int,0)
		*ans = append(*ans,newLevel)
	}
	(*ans)[length] = append((*ans)[length],point.Val)
	orderBottom(point.Left, length+1, ans)
	orderBottom(point.Right, length+1, ans)
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ans [][]int
	orderBottom(root,0,&ans)
	if len(ans)>0 {
		for i:=0;i<len(ans)/2 ;i++{
			ans[i],ans[len(ans)-i-1] = ans[len(ans)-i-1],ans[i]
		}
	}
	return ans
}

func main() {
	levelOrderBottom(nil)
}