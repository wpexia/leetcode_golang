package algorithm

import "math/rand"

// Quicksort 快速排序
func Quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left, right := 0, len(nums)-1
	pivot := rand.Int() % len(nums)
	nums[pivot], nums[right] = nums[right], nums[pivot]
	for i := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	Quicksort(nums[:left])
	Quicksort(nums[left+1:])
	return nums
}
