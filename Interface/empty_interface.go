/*
	空接⼝可以表示任何类型
	通过断⾔来将空接⼝转换为指定类型
 	v, ok := p.(int) //ok=true 时为转换成功

	Go 接⼝最佳实践

	一、倾向于使⽤⼩的接⼝定义，很多接⼝只包含⼀个⽅法
		Example:

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

	二、较⼤的接⼝定义，可以由多个⼩接⼝定义组合⽽成
		Example:

type ReadWriter interface {
	Reader
	Writer
}

	三、只依赖于必要功能的最⼩接⼝
		Example:

func StoredData(reader Reader) error {
 …
}

*/

package main

import (
	"fmt"
)

func IfDoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
	} else if i, ok := p.(string); ok {
		fmt.Println("String", i)
	} else {
		fmt.Println("Unknown", i)
	}
}

func SwitchDoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown", v)
	}
}

func main() {
	IfDoSomething("hello")
	SwitchDoSomething(100)
}
