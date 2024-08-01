package main

import (
	"fmt"
	"time"
)

func main() {
	var intSlice = []int{1, 2, 34, 6, 8, 89, 0}

	for _, va := range intSlice {
		go func(va int) {
			fmt.Println(va)
		}(va)
	}

	time.Sleep(2 * time.Second)
}
