package main

import "fmt"

// 插入排序函数
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i] //[9,5,1,4,3]
		j := i - 1

		// 将key与已排序元素从后向前比较，并将比key大的元素后移
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		// 将key插入到正确的位置
		arr[j+1] = key
	}
}

func main() {
	arr := []int{9, 5, 1, 4, 3}
	fmt.Println("Original array:", arr)
	insertionSort(arr)
	fmt.Println("Sorted array:  ", arr)
}
