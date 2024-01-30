package main

import (
	"fmt"
	"time"
)

func main() {
	mytick := time.NewTicker(time.Second * 2)
	defer mytick.Stop() //定时器不用了需要关闭
	done := make(chan struct{})
	go func() {
		for {
			time.Sleep(time.Second * 10)
			done <- struct{}{}
		}
	}()
	for {
		select {
		case <-done:
			fmt.Println("done!!!!")
			return
		case t := <-mytick.C:
			fmt.Printf("the curtime is %v\n", t)
		}
	}

}
