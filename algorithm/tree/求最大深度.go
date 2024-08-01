package main

import (
	"fmt"
	"math"
)

// TreeNode 是二叉树的节点结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// arrayToTree 将数组转换为二叉树
func arrayToTree(nums []int, index int) *TreeNode {
	if index >= len(nums) || nums[index] == -1 {
		return nil
	}

	root := &TreeNode{Val: nums[index]}
	root.Left = arrayToTree(nums, 2*index+1)
	root.Right = arrayToTree(nums, 2*index+2)

	return root
}

// maxDepth 计算二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// 返回左右子树深度的较大值加上当前节点的深度（1）
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}

func main() {
	// 示例输入数组 [3, 9, 20, null, null, 15, 7]
	inputArray := []int{3, 9, 20, -1, -1, 15, 7}

	// 将数组转换为二叉树
	root := arrayToTree(inputArray, 0)

	// 计算最大深度
	result := maxDepth(root)
	fmt.Println("二叉树的最大深度是:", result)
}
