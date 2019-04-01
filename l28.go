package main

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		flag := true
		for j := 0; j < len(needle); j++ {
			println(i, j)
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	next[0] = -1
	for i := 0; i < len(needle)-1; i++ {
		k := next[i]
		for k != -1 {
			if needle[i] == needle[k] {
				break
			} else {
				k = next[k]
			}
		}
		next[i+1] = k + 1
	}
	j := 0
	i := 0
	for i < len(haystack) && j < len(needle) {
		if j == -1 {
			j++
			i++
			continue
		}
		if haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return i - j
			} else {
				continue
			}
		} else {
			j = next[j]
		}
	}
	return -1
}

func main() {
	println(strStr2("AAAAA", "BBA"))
}
