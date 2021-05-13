package main

func maxProfit(k int,prices []int) int {
	lenPrices := len(prices)
	if lenPrices<2 {
		return 0
	}
	if k > lenPrices {
		k = lenPrices
	}
	lastResult := make([]int,k+1, k+1)
	tempMax := make([]int,k+1, k+1)
	maxResult:=0
	for i:= lenPrices-1;i>=0;i--{
		for j:=k;j>=0;j-- {
			if j == k{
				if prices[i]>tempMax[j] {
					tempMax[j] = prices[i]
				}
				continue
			}
			max := lastResult[j]
			if tempMax[j+1]-prices[i] > max{
				max = tempMax[j+1]-prices[i]
			}
			if max > maxResult {
				maxResult = max
			}
			if lastResult[j] + prices[i]>tempMax[j]{
				tempMax[j] = lastResult[j]+ prices[i]
			}
			if max>lastResult[j]{
				lastResult[j] = max
			}
		}
	}
	return maxResult
}

func main() {
	println(maxProfit(8,[]int{106,373,495,46,359,919,906,440,783,583,784,73,238,701,972,308,165,774,990,675,737}))
}