package main

import (
	"fmt"
	"math"
)

func main() {
	dividend := 10
	divisor := 3
	result := dividend / divisor

	// 保留两位小数
	roundedResult := math.Round(float64(result*100)) / 100
	fmt.Println("Result:", roundedResult)
}
