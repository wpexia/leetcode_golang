package main

//Find 二分查找
func Find(nums []int, left int, right int, target int) int {
	if left == right {
		if nums[left] == target {
			return left
		}
		return -1
	}
	pivot := (left + right) / 2
	if target <= nums[pivot] {
		return Find(nums, left, pivot, target)
	}
	return Find(nums, pivot+1, right, target)
}
