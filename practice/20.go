package main

import "fmt"

// 查询单链表有没有环
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := &ListNode{Val: 1}
	//node2 := &ListNode{Val: 2}
	//node3 := &ListNode{Val: 3}
	//node4 := &ListNode{Val: 4}

	//node1.Next = node2
	//node2.Next = node3
	//node3.Next = node4
	node1.Next = nil
	fmt.Println(findRing(node1))
}

func findRing(node *ListNode) bool {
	if node == nil || node.Next == nil {
		return false
	}

	slow, fast := node, node.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
