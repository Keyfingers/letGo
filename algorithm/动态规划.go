/*
动态规划（Dynamic Programming，简称DP）是一种算法策略，它适用于具有重叠子问题和最优子结构特性的问题。
动态规划通常用于解决多阶段决策过程的最优化问题，通过存储中间结果来避免重复计算，从而提高计算效率。

动态规划的关键特性：
重叠子问题：问题可以分解为多个子问题，这些子问题会重复出现。动态规划通过存储这些子问题的解，避免重复计算。
最优子结构：问题的最优解包含其子问题的最优解。这意味着，通过找到所有子问题的最优解，可以构建出原问题的最优解。

何时使用动态规划：
	问题规模较大：当问题规模较大，且存在大量重复计算时，动态规划可以显著减少计算量。

	优化问题：需要找到最大值、最小值或其他最优化条件的问题，如求最大子数组和、最短路径等。

	计数问题：需要计算满足特定条件的方案数，如切割问题、组合问题等。

	具有明确阶段的问题：问题可以自然地分解为一系列决策阶段，每个阶段都有若干选择，如棋类游戏、资源分配问题等。

	复杂度很高：问题的递归解法或简单迭代解法的时间复杂度过高，不适用于实际问题的解决。
*/

package main

import "fmt"

// knapsack 使用动态规划解决0/1背包问题
func knapsack(capacity int, weights []int, values []int, n int) int {
	// 初始化动态规划表格
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// 动态规划填表过程
	for i := 1; i <= n; i++ {
		for j := 0; j <= capacity; j++ {
			if weights[i-1] <= j {
				// 如果物品i可以放入背包，考虑放入和不放入两种情况
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			} else {
				// 如果物品i不能放入背包，则不改变当前背包的价值
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// 返回背包能够装载的最大价值
	return dp[n][capacity]
}

// max 返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 物品的重量
	weights := []int{2, 3, 4, 5}
	// 物品的价值
	values := []int{3, 4, 5, 6}
	// 背包的容量
	capacity := 5
	// 物品的数量
	n := len(values)

	// 计算最大价值
	maxVal := knapsack(capacity, weights, values, n)
	fmt.Printf("背包能够装载的最大价值为: %d\n", maxVal)
}
