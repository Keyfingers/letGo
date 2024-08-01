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

	for {
		select {
		case achenl <- true:
			fmt.Print("a")
			bchnel <- true
		}
	}

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		for {
			select {
			case bchnel <- true:
				defer wg.Done()
				fmt.Print("b")
				achenl <- true
			}
		}
	}(&wg)

	wg.Wait()

	/*
		1、打印代码
		2、网络协议
		3、逻辑问题
	*/

}
