package main

func maxProfit(prices []int) int {
	if len(prices) <2 {
		return 0
	}
	tempMax := 0
	lastResult := 0
	for i:= len(prices) - 1;i>=0;i-- {
		if tempMax - prices[i] > lastResult {
			if prices[i]+lastResult > tempMax{
				lastResult,tempMax = tempMax - prices[i],prices[i]+lastResult
			} else{
				lastResult = tempMax - prices[i]
			}
		} else if prices[i]+lastResult > tempMax {
			tempMax = prices[i]+lastResult
		}
	}
	return lastResult
}


func main() {
	println(maxProfit([]int{7,1,5,3,6,4}))
}
