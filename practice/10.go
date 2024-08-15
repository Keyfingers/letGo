package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int, target int) [][]int {
	sort.Ints(nums) // 对数组进行排序
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 跳过重复的元素
			continue
		}
		left, right := i+1, len(nums)-1

		for left < right {
			total := nums[i] + nums[left] + nums[right]
			if total == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// 跳过重复的元素
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if total < target {
				left++ // 需要增加和
			} else {
				right-- // 需要减小和
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	target := 0
	fmt.Println(threeSum(nums, target))
}
