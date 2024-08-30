package main

import "fmt"

func main() {
	var arr = []int{1, -2, 3, 10, -4, 7, 2, -5}

	var left, right, resl, resr = 0, 0, 0, 0
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	maxSum := dp[0]

	for i := 1; i < len(arr); i++ {
		right++
		//取最大和
		dp[i] = max1(dp[i-1]+arr[i], arr[i])

		if dp[i-1]+arr[i] < arr[i] { //当前值比(前一个和值和当前值)加起来还大，那么子序列在后面，做指针右移
			left = right
		}
		if dp[i] > maxSum || (dp[i] == maxSum && (right-left+1) > (resr-resl+1)) {
			maxSum = dp[i]
			resl = left
			resr = right
		}
	}
	//取值
	res := make([]int, (resr - resl + 1))
	for i := resl; i <= resr; i++ {
		res[i-resl] = arr[i]
	}
	fmt.Println(res)
}

func max1(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
