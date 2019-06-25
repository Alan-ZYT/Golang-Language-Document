package main

import "fmt"

func main() {
	/*
		go语言的数据类型：
			基本数据类型：
					int，float，bool，string

			复合数据类型：
					array，slice，map，function，pointer，struct，interface ...

		函数的类型：
				func(参数列表的数据类型)(返回值列表的数据类型)
	*/

	a := 10
	fmt.Printf("%T\n", a) //int

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", b) //[5]int
	/*
		对数组而言:数组的长度不同,存储的数据类型不同,导致数组自己的类型也不同
			example:
					[5]string
					[6]float64
	*/

	c := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", c) //[]int

	d := make(map[int]string)
	fmt.Printf("%T\n", d)
	/*
		对于Map而言: 所存储的键值类型不同,导致Map的类型也不同
		example:
				map[string]string
				map[string]map[int]string
	*/

	//获取函数数据类型
	//注意:fmt.Printf("%T\n", fun1) fun1后面不能加()例如:fun1(),加()代表函数的调用,不加()代表这个函数的本身
	fmt.Printf("%T\n", fun1) //func()
	fmt.Printf("%T\n", fun2) //func(int) int
	fmt.Printf("%T\n", fun3) //func(float64, int, int) (int, int)
	fmt.Printf("%T\n", fun4) //func(string,string,int,int)(string,int ,float64)
}

func fun1() {}

func fun2(a int) int {
	return 0
}

func fun3(a float64, b, c int) (int, int) {
	return 0, 0
}

func fun4(a, b string, c, d int) (string, int, float64) {
	return "", 0, 0
}
