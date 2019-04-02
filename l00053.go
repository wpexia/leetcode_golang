package main

func maxSubArray(nums []int) int {
	temp := 0
	max := -2147483648
	for i := 0; i < len(nums); i++ {
		if temp+nums[i] > max {
			max = temp + nums[i]
		}
		if temp+nums[i] > 0 {
			temp = temp + nums[i]
		} else {
			temp = 0
		}
	}
	return max
}
