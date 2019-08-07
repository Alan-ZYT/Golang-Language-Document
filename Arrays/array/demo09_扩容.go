package main

import "fmt"

func main() {

	/*
		切片叫做变长数组，长度不固定
			len(),cap()

		当向切片中添加数据时，如果没有超过容量，直接添加，如果超过容量，自动扩容(成倍增长)
		3-->6-->12-->24-->48
		4-->8-->16-->32
	*/

	s1 := []int{1, 2, 3}
	fmt.Println("len:", len(s1), "cap:", cap(s1), s1) // len=3, cap=3
	fmt.Printf("%p\n", s1)                            //将s1中存储的地址进行打印输出

	s1 = append(s1, 4, 5)
	fmt.Println("len:", len(s1), "cap:", cap(s1), s1) // len=5, cap=6
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 6, 7, 8)
	fmt.Println("len:", len(s1), "cap:", cap(s1), s1) // len=8, cap=12
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 9, 10)
	fmt.Println("len:", len(s1), "cap:", cap(s1), s1) // len=8, cap=12
	fmt.Printf("%p\n", s1)
	s1 = append(s1, 11, 12, 13, 14)
	fmt.Println("len:", len(s1), "cap:", cap(s1), s1) // len=8, cap=12

	arr := []int{100, 200, 300, 400}
	fmt.Println(s1)
	fmt.Println(arr)
	s1 = append(s1, arr...)

}
