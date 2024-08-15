package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg.Add(len(arr))
	for _, va := range arr {
		go func(va int) {
			fmt.Println(va)
			wg.Done()
		}(va)
	}
	wg.Wait()
}
