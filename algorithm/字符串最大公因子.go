package main

//
///*
//对于字符串 s 和 t，只有在 s = t + ... + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。
//
//给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。
//
//
//示例 1：
//
//输入：str1 = "ABCABC", str2 = "ABC"
//输出："ABC"
//示例 2：
//
//输入：str1 = "ABABAB", str2 = "ABAB"
//输出："AB"
//示例 3：
//
//输入：str1 = "LEET", str2 = "CODE"
//输出：""
//*/
//
//import (
//	"fmt"
//)
//
//// 辗转相除法求最大公因子
//func gcd(a, b int) int {
//	for b != 0 {
//		a, b = b, a%b
//	}
//	return a
//}
//
//func gcdOfStrings(str1 string, str2 string) string {
//	// 如果两个字符串的长度不同，则最长公因子肯定为空
//	if len(str1) != len(str2) {
//		return ""
//	}
//
//	// 计算两个字符串的最大公因子长度
//	gcdLen := gcd(len(str1), len(str2))
//
//	// 获取最长公因子
//	divisor := str1[:gcdLen]
//
//	// 检查最长公因子是否同时整除两个字符串
//	if len(str1)/gcdLen == len(str1)/gcdLen && len(str2)/gcdLen == len(str2)/gcdLen &&
//		divisor*(len(str1)/gcdLen) == str1 && divisor*(len(str2)/gcdLen) == str2 {
//		return divisor
//	} else {
//		return ""
//	}
//}
//
//func main() {
//	fmt.Println(gcdOfStrings("ABCABC", "ABC"))  // Output: "ABC"
//	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // Output: "AB"
//	fmt.Println(gcdOfStrings("LEET", "CODE"))   // Output: ""
//}
