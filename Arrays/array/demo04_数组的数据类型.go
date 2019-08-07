package main

import "fmt"

func main() {
	/*
		数据类型：[size]type
			数据类型：
				基本类型：int，float，string，bool
				复合类型：array，slice。。。

			值类型：理解为存储的数值本身。
				将数据传递给其他的变量，传递的是数据的副本(备份)
				int,float,string,bool,array

			引用类型：理解为存储的数据的内存地址。
				slice
	*/
	num := 10
	fmt.Printf("%T,%d\n", num, num)
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [3]float64{2.15, 3.14, 6.19}
	arr3 := [4]int{5, 6, 7, 8}
	arr4 := [2]string{"hello", "world"}
	fmt.Printf("arr1:%T,%v\n", arr1, arr1) //[4]int
	fmt.Printf("arr2:%T\n", arr2)          //[3]float64
	fmt.Printf("arr3:%T\n", arr3)          //[4]int
	fmt.Printf("arr4:%T\n", arr4)          //[2]string

	fmt.Println("--------------------------------")
	num2 := num                             //值传递
	fmt.Println("num,", num, "num2,", num2) //10 10
	num2 = 20
	fmt.Println("num,", num, "num2,", num2)

	//数组呢？
	arr5 := arr1
	fmt.Println(arr1, arr5) //[1 2 3 4] [1 2 3 4]
	arr5[0] = 100           //修改数组的第一个元素
	fmt.Println(arr1, arr5) //[1 2 3 4] [100 2 3 4]

	//==,比较数值是否相等
	a := 3                    //int
	b := 4                    //int
	fmt.Println(a == b)       //比较a和b的数值是否相等
	fmt.Println(arr5 == arr1) //比较数组的对应的数值是否相等 [4]int,[4]int
	//fmt.Println(arr1 == arr2) //(mismatched types [4]int and [3]float64)
}
