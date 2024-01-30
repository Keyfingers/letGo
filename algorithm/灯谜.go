package main

var a = 66
var sum = 1

func main() {
	for i := 1; i <= 99; i++ {
		temp := a - i
		sum = temp * sum
	}

	println(sum)
}
