package main

import "fmt"

func main() {
	numsChan := make(chan int, 10)

	//写
	go func() {
		for i := 0; i < 50; i++ {
			numsChan <- i
			fmt.Println("写入数据", i)
		}
		fmt.Println("数据全部写完毕，准备关闭管道")
		close(numsChan)
	}()

	//读
	//直接这样for range读取会报错的，因为主协程不知道管道是否已经写完，所以会一直在这里等待
	//怎么解决呢，需要在写入端将管道关闭close(numsChan)
	for v := range numsChan {
		fmt.Println("读取数据", v)
	}

	fmt.Println("OVER！")

}
