package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//注意端口号前面的冒号不能省略 ：
	conn, err := net.Dial("tcp", ":8848")

	if err != nil {
		fmt.Println("net dail err:", err)
		return
	}

	fmt.Println("client与server连接建立成功！")

	sendData := []byte("hello")

	for {
		//向服务器发送数据
		cnt, err := conn.Write(sendData)
		if err != nil {
			fmt.Println("conn write err:", err)
			return
		}

		fmt.Println("client =====> server 成功 cnt:", cnt, ",data:", string(sendData))

		//接收服务器返回的数据
		//创建一个容器，用于接收服务器返回的数据
		buf := make([]byte, 1024)
		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn read err:", err)
			return
		}

		fmt.Println("client <==== 接收 server，cnt:", cnt, ",data:", string(buf[:cnt]))
		time.Sleep(1 * time.Second)
	}
}
