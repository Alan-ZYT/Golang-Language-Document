package main

import "fmt"

func main() {

	/*
		因为切片是引用类型的数据，直接拷贝的是地址。
		浅拷贝：拷贝的是数据地址。
			导致多个变量指向同一块内存
			引用了诶行的数据，默认都是浅拷贝：slice，map
		深拷贝：拷贝的是数据本身
			值类型的数据，默认都是深拷贝：array，int，float，string，bool，struct

		C语言：
			step1：使用内存，自己开辟，创建变量，数组，
			step2：使用
			step3：手动归还

		Java，Python，Go。。。语言：垃圾自动回收：GC(garbage Collect)

			gc(),通知GC来收垃圾。

	*/

	//切片的深拷贝
	s1 := []int{1, 2, 3, 4}

	s2 := make([]int, 0, 0) //
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1, s2)

	s2[0] = 100
	fmt.Println(s1, s2)

	//copy()
	s3 := []int{7, 8, 9}
	fmt.Println(s1)
	fmt.Println(s3)
	//copy(s1,s3) //将s3中的元素，拷贝到s1中
	//copy(s3,s1)//将s1中的元素，拷贝到s3中

	copy(s3[1:], s1[2:])
	fmt.Println(s1)
	fmt.Println(s3)

}
