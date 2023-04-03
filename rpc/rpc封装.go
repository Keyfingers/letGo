package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

//要求服务端在注册rpc对象时，能编译检测出 注册对象是否合法

// 在接口中定义方法原型
type MyInterface interface {
	HelloWorld(string, *string) error
}

func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

type MyClient struct {
	c *rpc.Client
}

func InitClient(addr string) MyClient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return MyClient{
		c: conn,
	}
}

func (this *MyClient) HelloWorld(a string, b *string) error {
	return this.c.Call("hello.HelloWorld", a, b)
}
