package main


func maxProfit(prices []int) int {
	lenPrices := len(prices)
	if lenPrices<2 {
		return 0
	}
	result := make([][]int,lenPrices+1)
	result[lenPrices] = []int{0,0,0}
	result[lenPrices-1] = []int{0,0,0}
	maxResult:=0
	for i:= lenPrices-2;i>=0;i--{
		result[i] = []int{0,0,0}
		for j:=0;j<2;j++ {
			max := result[i+1][j]
			for k:=i+1;k<lenPrices;k++{
				if result[k+1][j+1]+prices[k]-prices[i]>max {
					max = result[k+1][j+1]+prices[k]-prices[i]
				}
			}
			result[i][j] = max
			if max> maxResult {
				maxResult = max
			}
		}
	}
	return maxResult
}


func main(){
	println(maxProfit([]int{1,2,3,4,5}))
}