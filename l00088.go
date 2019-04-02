package main

func myPrint(nums1 []int) {
	print("nums:  ")
	for _, item := range nums1 {
		print(item, " ")
	}
	println()
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	point := m - 1
	for i := n - 1; i >= 0; {
		if point < 0 || nums2[i] > nums1[point] {
			for j := m + n - i - 1; j > point+1; j-- {
				nums1[j] = nums1[j-1]
			}
			nums1[point+1] = nums2[i]
			i--
		} else {
			point--
		}
	}
}

func main() {
	merge([]int{2, 0}, 1, []int{1}, 1)
}
