package main

import "fmt"

func main() {

	/*
		Golang中new和make的区别
		区别在于new分配内存，make初始化slice，map和channel类型
		https://www.jianshu.com/p/c03f703cd738
	*/
	s1 := new([]int)
	fmt.Println(s1, len(*s1), cap(*s1)) // &[] 0 0

	s2 := make([]int, 5, 50)
	fmt.Println(s2, len(s2), cap(s2)) // [0 0 0 0 0] 5 50
}
