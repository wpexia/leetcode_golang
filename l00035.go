package main

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		} else {
			if nums[i] > target {
				return i
			}
		}
	}
	return len(nums)
}
func main() {
	println(searchInsert([]int{1, 3, 5, 6}, 2))
}
