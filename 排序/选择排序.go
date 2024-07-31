package main

import "fmt"

/*
选择排序是一种简单直观的排序算法，无论什么数据进去都是 O(n²) 的时间复杂度。
所以用到它的时候，数据规模越小越好。唯一的好处可能就是不占用额外的内存空间了吧。
*/

/*
首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
重复第二步，直到所有元素均排序完毕。
*/

func selectionSort(arr []int) []int {
	l := len(arr)
	if l == 0 {
		return arr
	}
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sorts := selectionSort(array)
	fmt.Println(sorts)
}
