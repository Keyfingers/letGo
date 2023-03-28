package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	//http包
	client := http.Client{}

	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("client get err")
		return
	}

	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")
	url := resp.Request.URL
	status := resp.Status
	code := resp.StatusCode
	body := resp.Body
	bytes, _ := io.ReadAll(body)
	all, err := ioutil.ReadAll(body)
	fmt.Println("body:", bytes)
	fmt.Println(all)
	fmt.Println(date)
	fmt.Println(server)
	fmt.Println(url)
	fmt.Println(status)
	fmt.Println(code)
}
