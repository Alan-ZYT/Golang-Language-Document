package main

import (
	"encoding/json"
	"fmt"
)

/*
	结构体反序列化:
	首先，依然需要创建结构体类型。要求该类型必须与之前序列化时指定的结构体类型完全一致（个数、类型、顺序）。

	序列化函数：
	import "encoding/json"
	func Unmarshal(data []byte, v interface{}) error

	参数1：待解析的 json编码字符串
	参数2：解析后传出的结果
	返回值：err


	Map反序列化:

	不需要使用make函数给m初始化，开辟空间。这是因为，在反序列化函数 Unmarshal() 中会判断传入的参数2，
	如果是map类型数据，会自动开辟空间。相当于是Unmarshal() 函数可以帮助我们做make操作。
	但！传参时需要注意，Unmarshal的第二个参数，是用作传出，返回结果的。因此必须传m的地址值。

	Slice反序列化:
	实现思路与前面两种的实现完全一致

*/

//结构体反序列化
//定义结构体
type Teacher struct {
	Name string
	Id   int
	Age  int
	Addr string
}

// 封装函数测试结构体反序列化
func structDeserial() {

	// 准备待反序列化的json字符串
	str := `{"Name":"Zhang","Id":901,"Age":35,"Addr":"BJ"}`

	// 定义Teacher类型变量
	var t Teacher

	// 使用 encoding/json 包中的 Unmarshal() 函数，反序列化
	err := json.Unmarshal([]byte(str), &t)
	if err != nil {
		fmt.Println("json.Unmarshal error", err)
		return
	}
	fmt.Println(t)
}

// Map反序列化
// 封装函数测试map反序列化
func mapDeserial() {
	// 准备待反序列化的json字符串
	str := `{"addr":"BJ","age":"36","name":"Zhang","id":"901"}`

	//定义map变量，类型必须与之前序列化的类型完全一致。
	var m map[string]interface{}

	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}
	// 查看反序列化结果
	fmt.Println(m)
}

// slice反序列化
// 封装函数测试slice反序列化
func sliceDeserial() {
	// 准备反序列化的json字符串
	str := `[{"addr":"四川绵阳江油","age":61,"name":"李白"},{"addr":["河南巩县","今河南巩义"],"age":58,"name":"杜甫"}]`

	//定义slice变量
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}
	//打印反序列化结果
	fmt.Println(slice)
}

func main() {
	// 调用测试结构体反序列化函数
	structDeserial()

	// 调用测试map反序列化函数
	mapDeserial()

	// 调用测试slice反序列化函数
	sliceDeserial()
}
