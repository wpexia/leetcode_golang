package main

func plusOne(digits []int) []int {
	plus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if plus == 0 {
			break
		}
		digits[i] = plus + digits[i]
		plus = digits[i] / 10
		digits[i] = digits[i] % 10
	}
	if plus > 0 {
		digits = append([]int{plus}, digits...)
	}
	return digits
}
