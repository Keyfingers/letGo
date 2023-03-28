package main

import "fmt"

type Teacher struct {
	Id      int    `json:"-"`            //使用json编码时，这个字段不参与编码
	Name    string `json:"student_name"` //在json编码时，这个字段会编码成student_name
	Age     int    `json:"age,string"`   //在json编码时，将age转成string类型
	gender  string `json:"gender"`
	Address string `json:"address,omitempty"` //在json编码时，如果这个字段是空，那么忽略掉不参与编码
}

func main() {
	stu := Teacher{
		Id:      10002,
		Name:    "zhangsan",
		Age:     18,
		gender:  "male",
		Address: "",
	}
	fmt.Println(stu)
}
