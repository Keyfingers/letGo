package main

import (
	"fmt"
)

func findGreatestSumOfSubArray(array []int) []int {
	// 特殊情况处理
	if len(array) == 0 {
		return nil
	}

	// 初始化dp数组和maxsum
	dp := make([]int, len(array))
	dp[0] = array[0]
	maxSum := dp[0]

	// 初始化滑动区间和结果区间
	left, right := 0, 0
	resl, resr := 0, 0

	for i := 1; i < len(array); i++ {
		right++
		// 状态转移：连续子数组和最大值
		dp[i] = max(dp[i-1]+array[i], array[i])

		// 如果当前值加上前一个值小于当前值，则更新滑动区间的左边界
		if dp[i-1]+array[i] < array[i] {
			left = right
		}

		// 更新最大值和结果区间
		if dp[i] > maxSum || (dp[i] == maxSum && (right-left+1) > (resr-resl+1)) {
			maxSum = dp[i]
			resl = left
			resr = right
		}
	}

	// 根据结果区间resl和resr取数组
	res := make([]int, resr-resl+1)
	for i := resl; i <= resr; i++ {
		res[i-resl] = array[i]
	}
	return res
}

// max函数，返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	array := []int{1, -2, 3, 10, -4, 7, 2, -5}
	maxSubarray := findGreatestSumOfSubArray(array)
	fmt.Printf("最大子数组为: %v\n", maxSubarray)
}
