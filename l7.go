package main

// Reverse https://leetcode.com/problems/reverse-integer/
func Reverse(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = -1 * x
	}
	result := 0
	for x > 0 {
		result = result*10 + x%10
		x = x / 10
	}
	result *= flag
	if result > 1<<31-1 || result < -1*1<<31 {
		result = 0
	}
	return result
}
