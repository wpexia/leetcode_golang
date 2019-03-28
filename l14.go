package main

func longestCommonPrefix(strs []string) string {
	if len(strs) ==0 {
		return ""
	}
	i := 0
	for {
		var ch byte = 0
		endFlag := false
		for _, item := range strs {
			if i >= len(item) {
				endFlag = true
				break
			}
			if ch == 0 {
				ch = item[i]
			} else {
				if item[i] != ch {
					endFlag = true
					break
				}
			}
		}
		if endFlag {
			break
		}
		i += 1
	}
	return strs[0][0:i]
}

func main() {
	print(longestCommonPrefix([]string{"ab", "abb", "abc"}))
}