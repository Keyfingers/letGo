package main

import (
	"fmt"
	"time"
)

func main() {
	//单向读通道
	//var chanReadOnly <- chan int
	//单向写通道
	//var chanWriteOnly chan <- int

	//1、在主函数中创建一个双向通道 numChan
	numChan := make(chan int, 5)
	//2、将numChan传递给producer，负责生产
	go producter(numChan) //双向通道可以赋值给同类型单向通道，反之则不行
	//3、将numChan传递给consumer，负责消费
	go consumer(numChan)

	time.Sleep(2 * time.Second)
	fmt.Println("OVER!")
}

func producter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		//data := <- out 写管道不允许呦读取操作
		fmt.Println("====》向管道中写入数据：", i)
	}
}

// consumer ===>提供一个只读通道
func consumer(in <-chan int) {
	for v := range in {
		fmt.Println("从管道读取读取数据：", v)
	}
}
