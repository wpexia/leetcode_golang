package main

func isValid(s string) bool {
	maps := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune = make([]rune, len(s))
	point := -1
	flag := true
	for _, char := range s {
		if !flag {
			break
		}
		switch rune(char) {
		case ')', ']', '}':
			if point < 0 || stack[point] != maps[rune(char)] {
				flag = false
				break
			} else {
				point -= 1
			}
		default:
			point += 1
			stack[point] = rune(char)
		}
	}
	if point != -1 {
		flag = false
	}
	return flag
}

func main() {
	print(isValid("(){}[]"))
}
