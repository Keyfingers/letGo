package main

import "fmt"

func main() {
	arr := []int{5, 3, 1, 4, 7}
	ints := selectSorted(arr)
	fmt.Print(ints)
}

// 选择排序
func selectSorted(arr []int) []int {
	length := len(arr)
	if length == 0 {
		return arr
	}
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
