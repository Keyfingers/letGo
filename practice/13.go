package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var arr = []int{23, 32, 78, 43, 76, 65, 345, 762, 77, 85, 36, 44, 915, 86}
	timer := time.NewTimer(5 * time.Second)
	resultChan := make(chan bool)
	target := 345
	size := 3
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr) - 1
		}
		go findTarget(ctx, arr[i:end], target, resultChan)
	}
	select {
	case <-timer.C:
		fmt.Println("timeout")
		cancel()
	case <-resultChan:
		sprintf := fmt.Sprintf("found result=%d", target)
		fmt.Println(sprintf)
		cancel()
	}
	time.Sleep(time.Second * 2)
}

func findTarget(ctx context.Context, data []int, target int, result chan bool) {
	for _, datum := range data {
		select {
		case <-ctx.Done():
			fmt.Println("task canceled")
			return
		default:
		}
		if datum == target {
			result <- true
			return
		}
	}

}
