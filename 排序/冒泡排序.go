package main

import "fmt"

// 正序时最快，反序时最慢
func bubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sorts := bubbleSort(array)
	fmt.Println(sorts)
}
