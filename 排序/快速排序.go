package main

import "fmt"

func main() {
	arr := []int{15, 2, 3, 4, 5, 6, 7, 8, 9, 20, 11, 12, 13, 14, 1}
	ints := sortArray(arr)
	fmt.Print(ints)
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	// 选择数组中间元素作为基准值
	x := nums[(l+r)>>1]
	for i < j {
		for {
			i++
			if nums[i] >= x {
				break
			}
		}
		for {
			j--
			if nums[j] <= x {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}
