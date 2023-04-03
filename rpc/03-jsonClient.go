package main

import "fmt"

func main() {
	myClient := InitClient("127.0.0.1:8800")

	var resp string
	err := myClient.HelloWorld("杜甫", &resp)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(resp)
}
