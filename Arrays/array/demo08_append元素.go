package main

import (
	"fmt"
)

func main() {
	/*
		切片不同于数组，变长
			长度和容量不一定相同。

		append() 尾部追加元素
			slice = append(slice, elem1, elem2)
			slice = append(slice, anotherSlice...)
	*/

	s1 := make([]int, 0, 5) //len=0,cap=5
	fmt.Println(s1)
	s2 := make([]int, 3, 5) //len=3,cap=5
	fmt.Println(s2)
	s3 := []int{1, 2, 3} //len=3,cap=3
	fmt.Println(s3, len(s3), cap(s3))

	//操作切片中已有的元素，使用下标
	s2[0] = 1
	fmt.Println(s2)
	//向切片中添加元素
	s2 = append(s2, 10, 20, 30, 40) //向切片中添加元素，如果超过容量，自动扩容
	fmt.Println(s2)

	s2 = append(s2, s3...)
	fmt.Println(s2)
	//思考题：删除切片中某个元素？

}
