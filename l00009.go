package main

// isPalindrome https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == Reverse(x) {
		return true
	}
	return false
}
