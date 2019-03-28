package main
func removeDuplicates(nums []int) int {
 	if len(nums)==0 {return 0}
 	result := 1
 	for i:=1;i<len(nums);i++{
 		if nums[i] != nums[i-1]{
			nums[result] = nums[i]
 			result += 1
 		}
 	}
 	return result
}

func main() {
	print(removeDuplicates([]int{1,1,2}))
}