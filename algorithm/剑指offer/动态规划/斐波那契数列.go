package main

import "fmt"

func main() {
	fibonacci := Fibonacci(10)
	fmt.Print(fibonacci)
}
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	var (
		res = 0
		a   = 0
		b   = 1
	)
	for i := 1; i < n; i++ {
		//将子问题的答案保留
		res = a + b
		a = b
		b = res
	}
	return res
}
