package main

import "fmt"

func main() {
	/*
		函数指针：一个指针，指向了一个函数的指针。
			因为go语言中，function，默认看作一个指针，只不过是没有*号。[Go语法糖]

			引用类型数据: slice,map,function 存储的就是数据的地址

		指针函数：一个函数，该函数的返回值是一个指针。
	*/

	var a func()
	a = fun1 //a 也是函数类型
	a()      // fun1()......

	arr1 := fun2()
	fmt.Printf("arr1的类型：%T，地址：%p，数值：%v\n", arr1, &arr1, arr1) //arr1的类型：[4]int，地址：0xc000056140，数值：[1 2 3 4]

	arr2 := fun3()
	fmt.Printf("arr2的类型：%T，地址：%p，数值：%v\n", arr2, &arr2, arr2) //arr2的类型：*[4]int，地址：0xc000084020，数值：&[5 6 7 8]
	fmt.Printf("arr2指针中存储的数组的地址：%p\n", arr2)                  //arr2指针中存储的数组的地址：0xc0000561c0
}

func fun3() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("函数中arr的地址：%p\n", &arr) //函数中arr的地址：0xc0000561c0
	return &arr
}

func fun2() [4]int { //普通函数
	arr := [4]int{1, 2, 3, 4}
	return arr
}

func fun1() {
	fmt.Println("fun1()...")
}
