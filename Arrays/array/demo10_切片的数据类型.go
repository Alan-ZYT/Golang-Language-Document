package main

import "fmt"

func main() {
	/*
		值类型：数组，int，float，string，bool
			传递的是数据副本

		引用类型：slice
			传递的地址,多个变量指向了同一块内存地址
	*/

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1 //传递的是数据
	fmt.Println(arr1, arr2)
	arr2[0] = 100
	fmt.Println(arr1, arr2)

	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1)

	s2 := s1 //传递的地址
	fmt.Println(s2)
	s2[0] = 100
	fmt.Println(s1, s2)
}
