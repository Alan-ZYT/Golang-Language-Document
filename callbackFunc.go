package main

/*
	回调函数：
	本质： 函数指针

	函数指针定义语法：type 函数指针类型  指向的函数原型	example: type  FUNCP func(string, int) bool

	回调函数：用户自定义一个函数，不直接调用，当满足某一特定条件时，由函数指针完成调用 或者 由系统自动调用。

	指针变量大小和数据类型无关,和操作系统有关 64位操作系统:8byte 32位操作系统:4byte

	指针变量类型的作用:取数据值用,和大小没有关系

	输出结果:
	Type=main.FUNCP Size=8
	res = 11
	res2 = 0

	具体应用参考以下代码:

*/

import (
	"fmt"
	"unsafe"
)

//定义函数指针类型
type FUNCP func(x int, y bool) int //该函数指针指向一类函数,有一个int和一个bool做参数,返回int的这一类函数

func useCallback(x int, y bool, p FUNCP) int {
	return p(x, y)
}

//回调函数
func addOne(x int, y bool) int {
	if y == true {
		x += 1
	}
	return x
}

//回调函数
func subTen(x int, y bool) int {
	if y == true {
		x -= 10
	}
	return x
}

func main() {
	//直接调用函数
	//res:=addOne(10, true)
	//fmt.Println(res)

	//定义一个函数指针类型
	var p FUNCP

	p = addOne
	fmt.Printf("Type=%T\tSize=%d\n", p, unsafe.Sizeof(p))

	//使用回调函数
	res := useCallback(10, true, p)
	fmt.Println("res =", res)

	res2 := useCallback(10, true, subTen)
	fmt.Println("res2 =", res2)

}
