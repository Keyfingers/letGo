package main

import "fmt"

// 选择排序
func main() {
	array := []int{1, 2, 3, 4, 5}
	ints := selectSorted(array)
	fmt.Print(ints)
}

func selectSorted(arr []int) []int {
	l := len(arr)
	if l < 0 {
		return arr
	}
	for i := 0; i < l; i++ {
		max := i
		for j := i + 1; j < l; j++ {
			if arr[max] < arr[j] {
				max = j
			}
		}
		arr[i], arr[max] = arr[max], arr[i]
	}
	return arr
}
