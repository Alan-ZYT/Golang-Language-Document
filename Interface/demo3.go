/*
⾏为（⽅法）定义
*/
package main

import (
	"fmt"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

// 第⼀种定义⽅式在实例对应⽅法被调⽤时，实例的成员会进⾏值复制
//func (e Employee) String() string {
//	return fmt.Sprintf("Id:%s\tName:%s\tAge:%d", e.Id, e.Name, e.Age)
//}

// 通常情况下为了避免内存拷⻉我们使⽤第⼆种定义⽅式
func (e *Employee) String() string {
	return fmt.Sprintf("Id:%s\tName:%s\tAge:%d", e.Id, e.Name, e.Age)
}

func main() {
	fmt.Println("OOP")
	e := Employee{"1", "BOB", 28}
	fmt.Println(e.String())
}
