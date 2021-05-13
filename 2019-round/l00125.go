package main

func isUp(ch byte) bool{
	return ch >= byte('A') && ch <= byte('Z')
}

func isDown(ch byte) bool{
	return ch >= byte('a') && ch <= byte('z')
}

func isNum(ch byte) bool{
	return ch >= byte('0') && ch <= byte('9')
}

func isAplhaNumeric(ch byte) bool {
	return isUp(ch)||isDown(ch)||isNum(ch)
}

var UPDOWNDIFF byte = byte('a') - byte('A')

func isPalindrome(s string) bool {
	i:=0;
	j:=len(s)-1
	for i<j{
		if !isAplhaNumeric(s[i]){
			i++
			continue
		}
		if !isAplhaNumeric(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			if !(isUp(s[i]) && s[i]+ UPDOWNDIFF == s[j]) && !(isUp(s[j]) && s[j]+UPDOWNDIFF == s[i]) {
				return false
			}
		}
		i++
		j--
	}
	return true
}

func main(){
	println(byte('a'),byte('z'),byte('A'),byte('Z'))
	print(isPalindrome("0P"))
}