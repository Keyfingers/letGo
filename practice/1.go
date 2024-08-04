package main

import "fmt"

func main() {
	arr := []int{5, 2, 3, 4, 1}
	ints := maopaoSorted(arr)
	fmt.Print(ints)
}

func maopaoSorted(arr []int) []int {
	length := len(arr)
	if length < 0 {
		return arr
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
