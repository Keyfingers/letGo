package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	va := dfs(root, 0)
	fmt.Println(va)
}

func dfs(node *TreeNode, currentSum int) int {
	if node == nil {
		return 0
	}

	currentSum = currentSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return currentSum
	}
	//继续遍历左右子树
	return dfs(node.Left, currentSum) + dfs(node.Right, currentSum)
}
