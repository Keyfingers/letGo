package main

import (
	"fmt"
	"sync"
)

func main() {
	var achannel, bchannel = make(chan bool), make(chan bool)

	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case <-achannel:
				fmt.Print("a")
				bchannel <- true
			}
		}
	}()

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		var bcounter = 1
		for {
			select {
			case <-bchannel:
				if bcounter > 10 {
					wg.Done()
					return
				}
				fmt.Print("b")
				bcounter++
				achannel <- true

			}
		}
	}(&wg)

	//启动
	achannel <- true

	wg.Wait()
}
