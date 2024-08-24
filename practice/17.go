package main

import "fmt"

//动态规划思想，求连续子数组的最大和

func main() {
	var arr = []int{1, -2, 3, 10, -4, 7, 2, -5}
	subArraySum := MaxSubArraySum(arr)
	fmt.Print(subArraySum)
}

func MaxSubArraySum(array []int) []int {
	if len(array) <= 0 {
		return nil
	}
	dp := make([]int, len(array))
	dp[0] = array[0]
	maxSum := dp[0]
	left, right := 0, 0
	resl, resr := 0, 0
	for i := 1; i < len(array); i++ {
		right++
		//连续数组最大和
		dp[i] = Max(array[i], dp[i-1]+array[i])

		//如果当前值+前面的值，还不如当前值大，那么做指针右移
		if dp[i-1]+array[i] < array[i] {
			left = right
		}
		//更新最大值区间
		if dp[i] > maxSum || (dp[i] == maxSum && (right-left+1) > (resr-resl+1)) {
			maxSum = dp[i]
			resl = left
			resr = right
		}
	}
	res := make([]int, resr-resl+1)
	for i := resl; i <= resr; i++ {
		res[i-resl] = array[i]
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
