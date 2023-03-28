package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type World struct {
}

func (this *World) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好！"
	return nil
}
func main() {
	//1、注册rpc服务
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("注册rpc失败", err)
		return
	}
	//2、设置监听
	listen, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listen.Close()

	fmt.Println("开始监听.....")

	//3、建立连接
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("建立连接失败：", err)
		return
	}
	defer conn.Close()

	fmt.Println("建立连接成功....")

	//4、绑定服务
	rpc.ServeConn(conn)
}
