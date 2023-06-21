package main

import (
	"fmt"
	"math"
)

func main() {
	a := 6
	b := 3

	d := divide(a, b)

	fmt.Println(d)
	fmt.Println(math.MinInt32)

}

func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	var sign = 1
	if a > 0 && b < 0 || a < 0 && b > 0 {
		sign = -1
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	var res int64
	for {
		if a < b {
			break
		}
		// a< b ,res = 0
		var cur = 1
		var temp = b
		for temp+temp <= a {
			temp += temp
			cur += cur
		}
		res += int64(cur)
		a -= temp
	}
	return int(res * int64(sign))

}
