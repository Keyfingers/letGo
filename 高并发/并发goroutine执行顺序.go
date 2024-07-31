package main

import (
	"fmt"
	"sync"
)

// 使用通道通信来同步
func main() {
	ch := make(chan int) // 初始化通道
	var count int
	var wg sync.WaitGroup
	done := make(chan struct{})

	wg.Add(2)

	go func() {
		defer wg.Done()
		ch <- 1
		close(done) // 通知发送完毕
	}()

	go func() {
		defer wg.Done()
		<-done // 等待发送完毕的通知，否则第二个goroutine关闭了，第一个还在发，那就panic了（除了使用通道来同步，也可以使用 sync.Mutex 来同步）
		count++
		close(ch) // 确保关闭的是已初始化的通道
	}()

	<-ch
	wg.Wait()
	fmt.Println(count)
}

// 使用sync.mutex来同步  sync.Mutex 用于确保同一时刻只有一个 goroutine 执行关键部分代码
func function2() {
	ch := make(chan int) // 初始化通道
	var count int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	go func() { //第一个 goroutine 在向 ch 发送数据时，锁住 mu
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		ch <- 1
	}()

	go func() { //第二个 goroutine 在关闭 ch 通道时，也锁住 mu，确保两个操作不会同时进行
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		count++
		close(ch) // 确保关闭的是已初始化的通道
	}()

	<-ch
	wg.Wait()
	fmt.Println(count)
}
