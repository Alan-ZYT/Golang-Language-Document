/*
	[10]int 和 [20]int 是不同类型
	调用 printArray 函数会拷贝数组

	数组固定的长度，又是拷贝，影响性能，用起来不方便
	可以用一个指向数组的指针
	在Go语言中一般不直接使用数组
	使用切片

*/

package main

import "fmt"

// 数组是值传递
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr { // range 关键字获取数组下标和value
		fmt.Println(i, v)
	}
}

func main1() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5} // 类型不同 [3]int [5]int
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println("array definitions:")
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}

// 指向数组的指针
func ptrArray(ptrarr *[5]int) {
	//(*ptrarr)[0] = 100
	ptrarr[0] = 100 // 语法糖
	for _, v := range ptrarr {
		fmt.Println(v)
	}
}

func main2() {
	var arr1 [5]int
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(arr1)
	fmt.Println(arr3)

	fmt.Println("ptrArray arr1 and arr3...")
	ptrArray(&arr1)
	ptrArray(&arr3)
}
