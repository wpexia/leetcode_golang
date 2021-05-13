func getRow(rowIndex int) []int {
	result := make([]int,rowIndex+1)
	for i:=0;i<=rowIndex;i++ {
		result[0] = 1
		result[i] = 1
		for j:=i-1;j>0;j--{
			result[j] = result[j-1]+result[j]
		}
	}
	return result
}