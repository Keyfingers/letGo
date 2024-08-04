package main

import "fmt"

func main() {
	str := "abcdefg"
	s := reverseString(str)
	fmt.Print(s)
}

// 翻转字符串
func reverseString(s string) string {
	str := []rune(s)
	length := len(str)
	for i := 0; i < length/2; i++ {
		str[i], str[length-i-1] = str[length-i-1], str[i]
	}
	return string(str)
}
