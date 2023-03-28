package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//注册路由 func是回调函数，用于路由的响应。
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//request是用户的请求详情
		fmt.Println("request：", request)
		//writer是将数据返回客户端的
		io.WriteString(writer, "这是/name请求返回的数据")
	})

	fmt.Println("http server start！")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("http start failed!", err)
		return
	}

}
