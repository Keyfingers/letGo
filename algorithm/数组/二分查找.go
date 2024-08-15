package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 7, 9, 10}
	index := binarySearch(nums, 7)
	fmt.Println(index)
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		middle := left + ((right - left) >> 1)
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
