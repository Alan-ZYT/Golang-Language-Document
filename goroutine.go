package main

/*
	并发：go 关键字，启动一个goroutine，执行对应的函数。

	程序的执行过程：
		A：先创建主goroutine
		B：初始化操作
		C：执行main()
		D：main()结束了，主goroutine随之结束，程序结束。

	go语言的并发：go关键字
		系统自动创建并启动主goroutine，执行对应的main()
		用于自己创建并启动子goroutine，执行对应的函数

		go 函数() //go关键字创建并启动goroutine，然后执行对应的函数()，该函数执行结束，子goroutine也随之结束。

	子goroutine中执行的函数，往往没有返回值。如果有也会被舍弃。
*/

import (
	"fmt"
	"time"
)

func main() {

	//go关键字创建并启动goroutine
	go hello()
	go printNum()
	go printLetter()

	time.Sleep(time.Second)

	fmt.Println("main...over...")
}

func hello() {
	fmt.Println("hello...")
}

//定义打印100个数字函数
func printNum() {
	for i := 1; i <= 100; i++ {
		fmt.Println("子goroutine中数字i:", i)
		time.Sleep(1)
	}
}

//定义打印100个字母函数
func printLetter() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("\t子goroutine中字母：%d,%c\n", i, i)
		time.Sleep(1)
	}
}
