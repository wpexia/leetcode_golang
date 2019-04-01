func climbStairs(n int) int {
	temp1 := 0 
	temp2 := 1
	for i:=1;i<=n;i++ {
		temp1,temp2 = temp2,temp1 + temp2
	}
	return temp2
}