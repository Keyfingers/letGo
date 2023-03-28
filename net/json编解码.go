package main

import (
	"encoding/json"
	"fmt"
)

// 在网络传输中，把Student编码成json字符串
type Student struct {
	Id     int
	Name   string
	Age    int
	gender string
}

func main() {
	lily := Student{
		Id:     10001,
		Name:   "Lily",
		Age:    20,
		gender: "famale", //小写字母开头的，在json编码时会忽略掉
	}

	enCodeInfo, err := json.Marshal(&lily)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return
	}
	fmt.Println("enCodeInfo:", enCodeInfo)

	var lily2 Student
	if err := json.Unmarshal([]byte(enCodeInfo), &lily2); err != nil {
		fmt.Println("json unmarshal err")
		return
	}
	fmt.Println(lily2)

}
