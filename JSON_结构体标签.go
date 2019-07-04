package main

import (
	"encoding/json"
	"fmt"
)

/*
type Student struct {
	name string
	age  int
	id   int
	sex  string
}
*/

/*type Student struct {
	Name string `json:"-"`             //“-”作用是不进行序列化，效果和将结构体字段写成小写一样。
	Age  int    `json:"age,omitempty"` // “omitempty”作用是在序列化的时候忽略0值或空值。
	Id   int    `json:"idx,string"`    // 序列化时，类型转化为string
	Sex  string `json:"sex"`
}*/

type Student struct {
	Name string `json:"stu_name"`
	Age  int    `json:"stu_age"`
	Id   int
	Sex  string
}

func main() {

	stu := Student{
		Name: "alan",
		Age:  25,
		Id:   1,
		Sex:  "male",
	}

	data, err := json.MarshalIndent(stu, "\t", "")
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	fmt.Println(string(data))
}
