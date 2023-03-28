package main

import "fmt"

func main() {
	numchan := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			numchan <- i
			fmt.Println("写入数据", i)
		}

		close(numchan)
	}()

	for {
		v, ok := <-numchan
		if !ok {
			fmt.Println("管道已经关闭，准备退出")
			break
		}
		fmt.Println("v：", v)
	}
	fmt.Println("OVER!")
}
