package main

import "fmt"

func DoesNotContain(slice []int, val int) bool {
	set := make(map[int]bool)
	for _, item := range slice {
		set[item] = true
	}
	// 检查val是否在集合中
	return set[val]
}

func main() {
	var a = []int{1, 2, 3, 4}
	str := "23675482"
	contain := DoesNotContain(a, 5)
	fmt.Println(contain)
	fmt.Println(str[len(str)-1:])
}
