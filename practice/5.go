package main

import "fmt"

func main() {
	arr := []int{11, 2, 3, 4, 9, 6, 7, 8, 0, 10}
	ints := insertionSort(arr)
	fmt.Println(ints)
}

// 插入排序
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
