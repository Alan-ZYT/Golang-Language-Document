package main

import "fmt"

func main() {
	/*

		数据类型：
			基本类型：整数，浮点，布尔，字符串
			复合类型：array，slice，map，struct，function...

		数组：
			1.概念：存储一组相同数据类型的数据结构
				数组定长，不能改变长度，但是可以改变存储的数据。
				有序，有下标，可以操作下标获取数据
					数组存储单个元素。

			2.语法：
				创建数组
					var 数组名 [长度] 数据类型
					var 数组名 = [长度]数据类型{元素1,元素2。。}
					var 数组名 = [长度] 数据类型{index:元素，index:元素。。。}
					数组名 ：= [长度]数据类型{}
					数组名:=[...]数据类型{元素}，使用{}号内的元素的数量，推断数组的长度。。

				访问数组：

					进行赋值和取值，同变量类似的操作

					通过下标访问数组
						下标，理解为元素在数组中的位置，编号，索引，index。
							从0开始的整数，取值范围[0,len-1]
							如果超过范围，越界错误:invalid array index 5 (out of bounds for 4-element array)

			3.len和cap：go语言的内置函数
				len(数组/map/slice/string)，长度

				cap()，容量
	*/

	var num1 int
	num1 = 100
	fmt.Println(num1)

	num1 = 200
	fmt.Println(num1)

	//1.创建数组
	var arr [4]int
	//2.通过下标也叫索引，index，访问数组，可以进行赋值或取值
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	fmt.Println(arr[0]) //获取数组的第一个数
	fmt.Println(arr[3]) //获取数组的第4个数
	//fmt.Println(arr[5]) //invalid array index 5 (out of bounds for 4-element array)

	//3.遍历数组：依次访问数组中的每一个元素。
	//遍历打印数据
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	//fmt.Println(arr[0])
	//fmt.Println(arr[1])
	//fmt.Println(arr[2])
	//fmt.Println(arr[3])

	//4.数组的长度，容量
	fmt.Println("数组的长度：", len(arr)) //length获取长度：容器中实际上存储的数据量
	fmt.Println("数组的容量：", cap(arr)) //capicity,获取容量：容器中能够存储的最大的数量
	//因为数组是定长的容器，长度和容量是相同的。

	//5.数组定长的容器，大小无法改变，但是我们更改里面存储的数据
	arr[0] = 100
	fmt.Println(arr)

	//6.数组的其他创建方式
	var a [4]int //同 var a =[4]int{}
	fmt.Println(a)

	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	var c = [5]int{1, 2, 3}
	fmt.Println(c)

	var d = [5]int{1: 1, 3: 2}
	fmt.Println(d)

	e := [5]string{"rose", "王二狗", "jack"}
	fmt.Println(e)

	f := [...]int{1, 2, 3, 4, 5} //定义数组的时候使用...，根据{}中的数据，来推断该数组的长度
	fmt.Println(f, "长度：", len(f))

	g := [...]int{1: 3, 6: 5}
	fmt.Println(g, "长度：", len(g))
}
