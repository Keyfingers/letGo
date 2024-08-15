package main

import (
	"context"
	"fmt"
	"time"
)

/*
假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排序。限时5秒，使用多个goroutine查找切片中是否存在给定的值，在查找到目标值或者超时后立刻结束所有goroutine的执行。

比如，切片 [23,32,78,43,76,65,345,762,......915,86]，查找目标值为 345 ，如果切片中存在，则目标值输出"Found it!"并立即取消仍在执行查询任务的goroutine。

如果在超时时间未查到目标值程序，则输出"Timeout！Not Found"，同时立即取消仍在执行的查找任务的goroutine。
*/

func main() {
	timer := time.NewTimer(5 * time.Second)
	var arr = []int{23, 32, 78, 43, 76, 65, 345, 762, 77, 85, 36, 44, 915, 86}
	size := 3
	target := 345
	resultChan := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr) - 1
		}
		go findarr(target, ctx, arr, resultChan)
	}
	select {
	case <-timer.C:
		fmt.Println("timeout")
		cancel()
	case <-resultChan:
		fmt.Println("已找到")
		cancel()
	}
}
func findarr(target int, ctx context.Context, data []int, resultChan chan bool) {
	for _, va := range data {
		select {
		case <-ctx.Done():
			fmt.Println("task cancel")
			return
		default:

		}
		if va == target {
			resultChan <- true
			return
		}
	}
}
