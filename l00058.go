package main

func lengthOfLastWord(s string) int {
	result := 0
	lastResult := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == ' ' {
			result = 0
			continue
		} else {
			result++
			lastResult = result
		}
	}
	return lastResult
}
