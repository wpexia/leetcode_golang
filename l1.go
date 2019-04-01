package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := m[target-nums[i]]
		if ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	return []int{0, 0}
}
