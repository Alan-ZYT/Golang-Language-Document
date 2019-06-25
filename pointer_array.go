package main

import "fmt"

func main() {
	/*
		数组指针：首先是一个指针，一个数组的地址。
			*[4]Type

		指针数组：首先是一个数组，存储的数据类型是指针
			[4]*Type

			*[5]float64,指针，一个存储了5个浮点类型数据的数组的指针
			*[3]string，指针，数组的指针，存储了3个字符串
			[3]*string，数组，存储了3个字符串的指针地址的数组
			[5]*float64，数组，存储了5个浮点数据的地址的数组
			*[5]*float64，指针，一个数组的指针，存储了5个float类型的数据的指针地址的数组的指针
			*[3]*string，指针，存储了3个字符串的指针地址的数组的指针
			**[4]string，指针，存储了4个字符串数据的数组的指针的指针
			**[4]*string，指针，存储了4个字符串的指针地址的数组，的指针的指针
	*/

	//1.创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1) //[1 2 3 4]

	//2.创建一个指针，存储该数组的地址--->数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)         //&[1 2 3 4]
	fmt.Printf("%p\n", p1)  //数组arr1的地址 0xc000056140
	fmt.Printf("%p\n", &p1) //p1指针自己的地址 0xc000084020

	//3.根据数组指针，操作数组
	(*p1)[0] = 100
	fmt.Println(arr1) //[100 2 3 4]

	p1[0] = 200       //简化写法
	fmt.Println(arr1) //[200 2 3 4]

	//4.指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2) //[1 2 3 4]
	fmt.Println(arr3) //[0xc0000580c8 0xc0000580d0 0xc0000580d8 0xc0000580e0]
	arr2[0] = 100
	fmt.Println(arr2) //[100 2 3 4]
	fmt.Println(a)    // 1
	*arr3[0] = 200
	fmt.Println(arr3) //[0xc0000580c8 0xc0000580d0 0xc0000580d8 0xc0000580e0]
	fmt.Println(a)    //200

	b = 1000
	fmt.Println(arr2) //[100 2 3 4]
	fmt.Println(arr3) //[0xc0000580c8 0xc0000580d0 0xc0000580d8 0xc0000580e0]
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i]) //200 1000 3 4
	}
}

