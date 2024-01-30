package main

import (
	"fmt"
	"time"
)

func main() {
	var count int
	for {
		select {
		case <-time.Tick(time.Second * 1):
			fmt.Println("case1")
			count++
			fmt.Println("count--->", count)
		case <-time.Tick(time.Second * 2):
			fmt.Println("case2")
			count++
			fmt.Println("count--->", count)
		}
	}

}

/*
可见 case2 永远没有被执行到，问题就出在代码逻辑上，首先看time.Tick方法。
它每次都会创建一个新的定时器，随着 for 循环进行， select 始终监听两个新创建的定时器，老的定时器被抛弃掉了，也就不会去读取老定时器中的通道。
select可以同时监听多个通道，谁先到达就先读取谁，如果同时有多个通道有消息到达，那么会随机读取一个通道，其他的通道由于没有被读取，所以数据不会丢失，需要循环调用
select 来读取剩下的通道。

总结：
tick创建完成之后，不是马上有一个tick.第一个tick在你设置的多少秒之后才会进行创建
golang当中的定时器实质上是这个单项的管道
time.NewTicker会定时触发任务，当下一次执行到来而当前任务画面执行完，会等待当前任务执行完毕在进行下一次任务。
Ticker和Timer的不同之处是，Ticker时间到达之后不需要人为的调用Reset方法来重新设置时间
*/
