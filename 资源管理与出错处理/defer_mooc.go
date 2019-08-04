/*
	Go语言是通过defer调用实现资源管理的
	defer调用 确保调用在函数结束时发生
	【本来调用一句话是当场发生，但用了defer确保在函数调用结束时发生】
	defer里面相当于有一个栈，这个栈是先进后出的 FILO

	程序中加上defer的好处：
	不用担心程序中有 return，甚至中间有 panic 都可以

	何时使用defer？
	Open/Close
	Lock/Unlock
	PrintHeader/PrintFooter

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer1() {
	defer fmt.Println(1) //后执行
	defer fmt.Println(2) //先执行
	fmt.Println(3)
	/*	return
		panic("error occurred...")
		fmt.Println(4)
	*/
}

//参数在 defer 语句时计算
func tryDefer2() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		if i == 6 {
			println("printed too many...")
			return
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//println("file already exists...")
		fmt.Println("Eroor:", err)
		return
	}
	defer file.Close() //关闭文件资源
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

// 1,1,2,3,5,8,13...
//   a,b
//     a,b
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	//tryDefer1()
	tryDefer2()
	writeFile("fib.txt")
}
