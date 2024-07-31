package main

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c //先调用此方法，但c通道最开始是空的，程序会阻塞在此处，直到c <- 0被赋值，解除阻塞！
}
func main() {
	go f()
	c <- 0
	print(a)
}
