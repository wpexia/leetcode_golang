package main


func maxProfit(prices []int) int {
	lenPrices := len(prices)
	if lenPrices<2 {
		return 0
	}
	lastResult :=[3]int{0,0,0}
	tempMax := [3]int{0,0,0}
	maxResult:=0
	for i:= lenPrices-1;i>=0;i--{
		for j:=0;j<3;j++ {
			if j == 2{
				if prices[i]>tempMax[j] {
					tempMax[j] = prices[i]
				}
				continue
			}
			max := lastResult[j]
			if tempMax[j+1]-prices[i] > max{
				max = tempMax[j+1]-prices[i]
			}
			if max> maxResult {
				maxResult = max
			}
			if lastResult[j]+ prices[i]>tempMax[j]{
				tempMax[j] = lastResult[j]+ prices[i]
			}
			if max>lastResult[j]{
				lastResult[j] = max
			}
		}
	}
	return maxResult
}


func main(){
	println(maxProfit([]int{6,1,3,2,4,7}))
}