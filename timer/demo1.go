package main

import (
	"fmt"
	"time"
)

func main() {
	myTimer := time.NewTimer(time.Second * 3) //初始化定时器
	var i = 0
	for {
		select {
		case <-myTimer.C:
			i++
			fmt.Printf("the counter is%d\n", i)
			myTimer.Reset(time.Second * 3) //注意需要重新设置
		}
	}
	myTimer.Stop() //不在使用需要将其停止
}
