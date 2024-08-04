package main

import "fmt"

// 冒泡排序
func main() {
	arr := []int{5, 4, 3, 2, 1}
	ints := maopaoSorted(arr)
	fmt.Print(ints)
}

func maopaoSorted(arr []int) []int {
	l := len(arr)
	if l < 0 {
		return arr
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
