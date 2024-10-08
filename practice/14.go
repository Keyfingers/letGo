/*
求子数组的最大和
题目:输入一个整形数组，数组里有正数也有负数。 数组中连续的一个或多个整数组成一个子数组，每个子数组都有一个和。 求所有子数组的和的最大值。要求时间复杂度为 O(n)。
例如输入的数组为1, -2, 3, 10, -4, 7, 2, -5，和最大的子数组为3, 10, -4, 7, 2， 因此输出为该子数组的和18。
*/
package main

import "fmt"

func main() {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}
	fmt.Println("最大子数组和是:", maxSum(arr)) // 应该输出18
}

func maxSum(arr []int) int {
	maxCurrent := arr[0]
	maxGlobal := arr[0]
	for i := 1; i < len(arr); i++ {
		// 如果当前元素比加上当前元素的局部最大和还要大，则重新开始计算局部最大和
		maxCurrent = max(arr[i], maxCurrent+arr[i])
		maxGlobal = max(maxGlobal, maxCurrent)
	}
	return maxGlobal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
