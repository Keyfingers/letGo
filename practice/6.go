package main

import "fmt"

func main() {
	array := []int{3, 1, 6, 9, 0, 12, 50, 33, 100}
	ints := sortedArray(array)
	fmt.Print(ints)
}

func sortedArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l <= r {
		return
	}

	i, j := l-1, r+1
	x := nums[l+r>>1]
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
