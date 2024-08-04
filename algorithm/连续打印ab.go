package main

import (
	"fmt"
	"sync"
)

func main() {
	//a打印a，b打印b
	var achenl = make(chan bool)
	var bchnel = make(chan bool)

	wg := sync.WaitGroup{}

	go func() {
		for {
			select {
			case <-achenl:
				fmt.Print("a")
				bchnel <- true
			}
		}
	}()

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		counter := 1
		for {
			select {
			case <-bchnel:
				if counter > 10 {
					wg.Done()
					return
				}
				fmt.Print("b")
				counter++
				achenl <- true
			}
		}
	}(&wg)

	achenl <- true

	wg.Wait()

	/*
		1、打印代码
		2、网络协议
		3、逻辑问题
	*/

}
