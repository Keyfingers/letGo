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

	fmt.Println("监听中.....")

	//Accept() (conn,error)
	conn, err := listen.Accept()

	if err != nil {
		fmt.Println("listen accept err!")
		return
	}

	fmt.Println("连接建立成功!")

	//创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024) //使用make来创建字节切片，byte ==> uint8

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

	//关闭连接
	conn.Close()

}
