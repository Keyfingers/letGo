package main

import "fmt"

type ListNodes struct {
	Val  int
	Next *ListNodes
}

func main() {
	node1 := &ListNodes{Val: 1}
	node2 := &ListNodes{Val: 2}
	node3 := &ListNodes{Val: 3}
	node4 := &ListNodes{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node1
	fmt.Println(JudgeIsCycle(node1))
}

func JudgeIsCycle(head *ListNodes) bool {
	if head == nil || head.Next == nil {
		return false
	}
	/*判断是否有环的思路是什么呢？
	其实很简单，此处才有快慢指针的思想，如果有环的话，那么慢指针是可以追上快指针的
	*/
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
