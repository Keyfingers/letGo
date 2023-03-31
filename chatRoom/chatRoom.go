package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type User struct {
	Name string
	Id   string
	//消息管道
	Msg chan string
}

// 创建全局map，保存所有用户
var allUsers = make(map[string]User)

// 接收用户发来的消息
var message = make(chan string, 10)

func main() {
	//启动服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listener start err:", err)
		return
	}
	//启动全局唯一go程，负责监听message通道，写给所有用户
	go broadcast()

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
	fmt.Println("启动业务...")
	//客户端与服务器建立连接时候，会有ip和port，可用作user的id
	clientAddr := conn.RemoteAddr().String()
	newUser := User{
		Name: clientAddr,
		Id:   clientAddr,
		Msg:  make(chan string, 10), //注意需要分配空间，否则无法写入数据
	}
	//添加user到map
	allUsers[newUser.Id] = newUser
	isQuit := make(chan bool)
	//创建一个重置计时器，监听用户是活着的
	restTimer := make(chan bool)
	//启动go程监听退出信号
	go watch(&newUser, conn, isQuit, restTimer)
	//启动go程，负责将msg信息返回客户端
	go writeBackToClient(&newUser, conn)
	//向message写入数据
	loginInfo := fmt.Sprintf("[%s]:[%s] ===>上线了\n", newUser.Id, newUser.Name)
	message <- loginInfo
	for {
		//具体业务逻辑
		buf := make([]byte, 1024)

		//读取客户端发来的数据
		cnt, err := conn.Read(buf)
		if cnt == 0 {
			fmt.Println("客户端ctrl+c，准备退出！")
			isQuit <- true
			//用户退出，需要删除map，关闭链接
		}
		if err != nil {
			fmt.Println("conn read err:", err, "cnt:", cnt)
			return
		}
		fmt.Println("服务器接收客户端发来的数据：", string(buf[:cnt-1]), ",cnt:", cnt)

		//业务逻辑处理----开始--------
		//1、查询当前所有用户 who
		//	a、判断接收的数据是不是who
		userInput := string(buf[:cnt-1])
		if len(userInput) == 3 && userInput == "who" {
			//	b、遍历allUser，将id和name拼接成字符串返回
			var userInfos []string
			for _, user := range allUsers {
				userInfo := fmt.Sprintf("id:%s,name:%s", user.Id, user.Name)
				userInfos = append(userInfos, userInfo)
			}
			//最终写到管道中一定是一个字符串
			r := strings.Join(userInfos, "\n")
			newUser.Msg <- r
		} else if len(userInput) > 8 && userInput[:6] == "rename" {
			newUser.Name = strings.Split(userInput, "|")[1]
			allUsers[newUser.Id] = newUser
			//通知客户端修改成功
			newUser.Msg <- "rename successful \n"
		} else {
			//如果用户输入不是特殊命令，只需要写到广播通道常规转发即可
			message <- userInput
		}

		restTimer <- true

	}
}

// 向所有用户广播消息，启动一个全局go程
func broadcast() {
	fmt.Println("广播go程启动成功！")
	defer fmt.Println("broadcast程序退出")
	for {
		//1、从message读取数据
		info := <-message
		fmt.Println("message接收到消息：", info)

		//2、将数据写入到每一个用户的msg管道中
		for _, user := range allUsers {
			//如果msg是非缓冲的，程序会阻塞到这里。因为不知道程序什么时候执行完
			user.Msg <- info
		}
	}

}

func writeBackToClient(user *User, conn net.Conn) {
	fmt.Sprintf("user:%s 的go程正在监听自己的msg通道：\n", user.Name)
	for data := range user.Msg {
		fmt.Sprintf("user: %s 写回客户端的数据为：%s\n", user.Name, data)
		_, _ = conn.Write([]byte(data))
	}
}

// 新启动一个go程用来监听用户的退出操作
func watch(user *User, conn net.Conn, isQuit, restTimer <-chan bool) {
	defer fmt.Println("watch 监听退出！")
	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("%s exit already!\n", user.Name)
			delete(allUsers, user.Id)
			message <- logoutInfo
			conn.Close()
			return
		case <-time.After(10 * time.Second):
			logoutInfo := fmt.Sprintf("%s timeout exit already! \n", user.Name)
			delete(allUsers, user.Id)
			message <- logoutInfo
			conn.Close()
			return
		case <-restTimer:
			fmt.Println("重置计数器")

		}
	}
}
