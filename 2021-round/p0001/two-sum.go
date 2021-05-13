package p0001

func twoSum(nums []int, target int) []int {
	record := make(map[int]int,0)
	for index,value := range nums{
		if preIndex,ok :=record[target - value];ok{
			return []int{preIndex,index}
		}
		record[value]= index
	}
	panic("can't found result")
}
