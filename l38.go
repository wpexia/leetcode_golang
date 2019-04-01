package main

import "strconv"

func countAndSay(n int) string {
    result := "1"
    for i:=1;i<=n;i++{
		temp := ""
		last := result[0]
		lastNum := 1
		for j:=1;j<len(result);j++{
			if result[j]==last {
				lastNum ++
				continue
			} else {
				temp += strconv.Itoa(lastNum)+string(last)
				last = result[j]
				lastNum = 1
			}
		}
		temp += strconv.Itoa(lastNum)+string(last)
		result = temp
    }
    return result
}

func main() {
	println(countAndSay(3))
}