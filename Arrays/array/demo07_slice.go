package main

import "fmt"

func main() {
	/*
		切片slice：
			同数组类似，也叫做变长数组。
			是一个引用类型的容器，指向了一个底层数组。

		创建：
			s1 := [] 类型{}
			s2 := make([]类型，len，cap)

		内置函数：
			len(),cap(),make()

			make(),创建slice，map，chan
	*/

	arr := [4]int{1, 2, 3, 4} //定长
	fmt.Println(arr)
	//1.创建一个切片
	var s1 []int
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4} //变长
	fmt.Println(s2)
	fmt.Printf("%T,%T\n", arr, s2) //[4]int, []int

	s3 := make([]int, 3, 8) //第一个参数：切片的类型，第二个参数：len，第三个参数容量，如果省略第三个数，默认同第二个参数
	fmt.Println(s3)
	fmt.Println("切片的长度：", len(s3), "，切片的容量：", cap(s3))

	//2.切片的操作,通过下标访问切片中的元素
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Println(s3)
	//s3[3] = 4 //panic: runtime error: index out of range
	//fmt.Println(s3)
	fmt.Println(s3[0])
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	for i, v := range s3 {
		fmt.Println(i, v)
	}

	//3.根据已有的数组创建切片

	/*
	 slice := arr[start:end]
	 	切片中的数据：[start,end)
	 	arr[:end],从头到end
	 	arr[start:]从start到末尾

	 从已有的数组上，直接创建切片，该切片的底层数组就是当前的数组。
	 长度是从start到end切割的数据量。但是容量从start到数组的末尾。

	*/

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s4 := a[:5] //[start,end)
	fmt.Println(a)
	fmt.Println(s4, "长度：", len(s4), "容量：", cap(s4))
	s5 := a[3:8]
	fmt.Println(s5, "长度：", len(s5), "容量：", cap(s5))
	s6 := a[6:]
	fmt.Println(s6)

	//更改数组的数据
	a[4] = 100
	fmt.Println(a)
	fmt.Println(s4)
	fmt.Println(s5)

	s4[3] = 200
	fmt.Println(s4)
	fmt.Println(a)
	fmt.Println(s5)

	//向s4切片中添加元素
	fmt.Println("-------------------------------")
	fmt.Println(a)
	fmt.Println(s4) //len=5,cap=10
	fmt.Println(s5) //len=5,cap=7

	s4 = append(s4, 1, 1, 1, 1) //修改数组中的元素
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(a)
	fmt.Println(s5)

	s4 = append(s4, 2, 2, 2, 2)
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(a)

}
