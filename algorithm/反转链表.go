package main

import "fmt"

// ListNode 定义链表节点的结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 反转链表的函数
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	prev = nil
	curr := head

	for curr != nil {
		next := curr.Next // 保存下一个节点
		curr.Next = prev  // 将当前节点的Next指向前一个节点，实现反转
		prev = curr       // 前一个节点推进到当前节点
		curr = next       // 当前节点移动到下一个节点
	}

	return prev // 返回反转后链表的头节点
}

func main() {
	// 创建一个示例链表: 1 -> 2 -> 3 -> 4 -> 5
	var head = &ListNode{Val: 1}
	curr := head
	for i := 2; i <= 5; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
	}

	// 打印原始链表
	printList(head)

	// 反转链表
	head = reverseList(head)
	// 打印反转后的链表
	printList(head)
}

// printList 打印链表的辅助函数
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, "->")
		curr = curr.Next
	}
	fmt.Println("null")
}
