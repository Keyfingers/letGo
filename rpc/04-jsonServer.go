package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

type world struct {
}

func (this *world) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好！"
	return nil
	//return errors.New("未知错误！")
	//即使这样，只要错误不为空，就不会返回值！
}
func main() {
	//1、注册rpc服务
	//err := rpc.RegisterName("hello", new(world))
	//if err != nil {
	//	fmt.Println("注册rpc失败", err)
	//	return
	//}
	RegisterService(new(world))
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
	jsonrpc.ServeConn(conn)
}

// echo -e '{"method":"hello.HelloWorld","params":["李白"],"id":0}' | nc 127.0.0.1 8800 充当客户端
