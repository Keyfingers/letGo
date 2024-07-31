package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	go func() {
		ch = make(chan int, 1) //由于 ch 是在第一个 goroutine 中初始化的，而第二个 goroutine 可能在 ch 被初始化之前就已经开始运行并尝试读取 ch。这导致了数据竞争的问题。
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second) //如果第二个 goroutine 在 ch 初始化之前尝试从通道读取数据，会导致该 goroutine 挂起，等待通道的数据。
		<-ch
	}(ch) //即使第一个 goroutine 初始化了 ch，由于没有任何同步机制（例如 sync.Mutex 或 sync.WaitGroup），在多核处理器上运行时，第二个 goroutine 可能看不到第一个 goroutine 对 ch 的初始化。
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
		//打印 #goroutines: 2 （一个是第二个goroutine，一个是主goroutine）
	}
}
