package main

import (
	"fmt"
	"time"
)

func main() {
	f := func() {
		fmt.Println("the func is call after 3 second")
	}
	myTime := time.AfterFunc(time.Second*3, f)

	defer myTime.Stop() //定时器不用了需要关闭
	time.Sleep(time.Second * 4)
}
