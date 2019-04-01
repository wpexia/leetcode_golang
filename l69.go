package main

func mySqrt(x int) int {
	result := 0
	for true {
		if result*result > x {
			return result - 1
		} else {
			result += 1
		}
	}
	return result - 1
}

func mySqrt1(x int) int {
	if x < 2 {
		return x
	}
	left := 0
	right := x / 2
	for true {
		if left == right {
			return left
		}
		temp := (left + right + 1) / 2
		if temp*temp > x {
			right = temp - 1
		} else {
			left = temp
		}
	}
	return 0
}

func main() {
	print(mySqrt1(9))
}
