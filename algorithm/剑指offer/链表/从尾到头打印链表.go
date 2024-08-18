package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listNode := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	tailToHead := printListFromTailToHead(listNode)
	fmt.Println(tailToHead)
}

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	var ans []int
	if head != nil {
		ans = append(printListFromTailToHead(head.Next), head.Val) //将每个节点的值添加到切片
	}
	return ans
}
