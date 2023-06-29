package main

import "fmt"

func main() {
	fa := f()
	fmt.Println(fa())
	fmt.Println(fa())
	fmt.Println(fa())
}

func f() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
