package main

import (
	"fmt"
)

type TreeNodes struct {
	Val   int
	Left  *TreeNodes
	Right *TreeNodes
}

func printFromTopToBottom(root *TreeNodes, res *[]int) {
	if root == nil {
		return
	}

	queue := []*TreeNodes{root} // 使用队列存储待访问的节点

	for len(queue) > 0 {
		node := queue[0]  // 访问队列的第一个节点
		queue = queue[1:] // 从队列中移除已访问的节点

		if node != nil {
			*res = append(*res, node.Val) // 将节点值添加到结果数组
		}

		// 将左右子节点加入队列，如果它们存在
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func levelOrder(root *TreeNodes) []int {
	res := []int{}
	printFromTopToBottom(root, &res)
	return res
}

func main() {
	// 构建示例二叉树
	//     8
	//    / \
	//   6  10
	//  / \   \
	// 2   1  #
	root := &TreeNodes{Val: 8}
	root.Left = &TreeNodes{Val: 6}
	root.Right = &TreeNodes{Val: 10}
	root.Left.Left = &TreeNodes{Val: 2}
	root.Left.Right = &TreeNodes{Val: 1}

	// 打印从上到下的所有节点值
	fmt.Println(levelOrder(root))
}
