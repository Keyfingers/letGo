package main

import (
	"fmt"
	"sync"
)

func main() {
	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wg := new(sync.WaitGroup)
	wg.Add(9)
	for _, va := range sliceInt {
		go func(va int) {
			defer wg.Done()
			fmt.Println(va)
		}(va)
	}
	wg.Wait()
}
