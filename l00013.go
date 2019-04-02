package main

import "fmt"

func romanToInt(s string) int {
	values1 := map[string]int{
		"CM": 900,
		"CD": 400,
		"XC": 90,
		"XL": 40,
		"IX": 9,
		"IV": 4,
	}

	values2 := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	result := 0
	for i := 0; i < len(s); {
		if i < len(s)-1 {
			flag := false
			for key, value := range values1 {
				if s[i:i+2] == key {
					result += value
					flag = true
					break
				}
			}
			if flag {
				i += 2
				continue
			}
		}
		for key, value := range values2 {
			if s[i:i+1] == key {
				result += value
				break
			}
		}
		i += 1
	}
	return result
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
