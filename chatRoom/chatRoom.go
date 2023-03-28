package main

import (
	"fmt"
	"net"
)

func main() {
	//启动服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listener start err:", err)
		return
	}
	fmt.Println("服务器已启动！")

	for {
		fmt.Println("监听中.....")
		//设置监听，建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
			return
		}

		//处理业务逻辑
		go handler(conn)
	}

}

func handler(conn net.Conn) {
	for {
		fmt.Println("启动业务...")
		//todo
		buf := make([]byte, 1024)

		//读取客户端发来的数据
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read err:", err)
			return
		}
		fmt.Println("服务器接收客户端发来的数据：", string(buf[:cnt-1]), ",cnt:", cnt)
	}
}
