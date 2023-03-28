package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//创建监听
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)

	listen, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("net listen err!")
		return
	}

	/*
		1、我们想要server能够接收多个连接，======>主协程负责监听，子协程用于数据处理
		2、每个连接可以接收处理多轮数据请求
	*/

	for {
		fmt.Println("监听中.....")

		//Accept() (conn,error)
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("listen accept err!")
			return
		}

		fmt.Println("连接建立成功!")

		go handFunc(conn)

	}

}

func handFunc(conn net.Conn) {
	for { //这个for循环，保证每一个连接可以多次接收处理客户端的请求
		//创建一个容器，用于接收读取到的数据
		buf := make([]byte, 1024) //使用make来创建字节切片，byte ==> uint8

		fmt.Println("准备读取客户端发来的数据.......")
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read err:", err)
			return
		}

		fmt.Println("client ===> server,长度：", cnt, "数据：", string(buf[0:cnt]))

		//服务器对客户端请求进行响应，将数据转成答谢 "hello"  ==> "HELLO"
		upperData := strings.ToUpper(string(buf[0:cnt]))
		cnt, err = conn.Write([]byte(upperData))
		fmt.Println("client <=== server,长度：", cnt, "，数据：", upperData)
	}
}
