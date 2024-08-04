package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sumClosest := threeSumClosest(nums, 0)
	fmt.Print(sumClosest)
}

// 三数之和 利用双指针求解
func threeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	sum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			current := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(sum-target)) > math.Abs(float64(target-current)) {
				sum = current
			}
			if current < target {
				l++
			} else if current == target {
				return target
			} else {
				r--
			}
		}
	}
	return sum
}
