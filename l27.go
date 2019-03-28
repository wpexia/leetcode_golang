func removeElement(nums []int, val int) int {
 	if len(nums)==0 {return 0}
 	result := 0
 	for i:=0;i<len(nums);i++{
 		if nums[i] != val{
			nums[result] = nums[i]
 			result += 1
 		}
 	}
 	return result
}