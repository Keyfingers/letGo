package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	depth := DepthTreeNode(node)
	fmt.Println(depth)
}

func DepthTreeNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := DepthTreeNode(root.Left)
	right := DepthTreeNode(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
