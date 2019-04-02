package main

import "strconv"

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	result_it := 0
	temp := 0
	max_length := len(a) + 1
	if len(b)+1 > max_length {
		max_length = len(b) + 1
	}
	result := make([]int, max_length)
	for i >= 0 || j >= 0 {
		result[result_it] = temp
		if i >= 0 {
			result[result_it] += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			result[result_it] += int(b[j] - '0')
			j--
		}
		temp = result[result_it] / 2
		result[result_it] = result[result_it] % 2
		result_it++
	}
	if temp > 0 {
		result[result_it] = temp
		result_it++
	}
	resultStr := ""
	for i := result_it - 1; i >= 0; i-- {
		resultStr += strconv.Itoa(result[i])
	}
	return resultStr
}

func main() {
	println(addBinary("11", "1"))
}
