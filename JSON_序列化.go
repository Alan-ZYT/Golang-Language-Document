package main

import (
	"encoding/json"
	"fmt"
)

/*
	序列化函数：Marshal
	import "encoding/json"
	func Marshal(v interface{}) ([]byte, error)


	Go 语言标准库的"encoding/json"包还提供了另外一个方法：MarshalIndent。
	该方法的作用与Marshall作用相同，只是可以通过方法参数，设置前缀、缩进等，
	对Json 多了一些格式处理，打印出来比较好看
	example: json.MarshalIndent(stu, "\t", "")

*/

// 定义一个结构体
type Person struct {
	Name string
	Age  int
	Sex  string
}

// 封装 struct 序列化测试函数
func structSerial() {
	// 定义结构体变量
	p := &Person{
		Name: "alan",
		Age:  23,
		Sex:  "male",
	}

	// 使用 encoding/json 包中的 Marshal() 函数将结构体变量进行序列化
	//data, err := json.Marshal(p)
	data, err := json.MarshalIndent(p, "\t", "")
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	// 查看序列化后的 json 字符串
	fmt.Println(string(data))
}

// 封装 map 序列化测试函数
func mapSerial() {
	var m = make(map[string]interface{})

	m["name"] = "alan"
	m["age"] = 23
	m["Sex"] = "male"

	data, err := json.MarshalIndent(m, "\t", "")
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	fmt.Println(string(data))
}

// 封装 slice 序列化测试函数
func sliceSerial() {
	var m = make(map[string]interface{})
	m["name"] = "alan"
	m["age"] = 23
	m["Sex"] = "male"

	var s = []interface{}{}
	s = append(s, m)

	var m1 = make(map[string]interface{})
	m1["name"] = "lishy"
	m1["age"] = 25
	m1["Sex"] = "female"
	s = append(s, m1)

	data, err := json.MarshalIndent(s, "\t", "")
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	// 调用测试结构体序列函数
	structSerial()

	// 调用测试map序列函数
	mapSerial()

	// 调用测试slice序列函数
	sliceSerial()
}
