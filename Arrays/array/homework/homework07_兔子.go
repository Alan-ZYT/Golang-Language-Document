package main

import "fmt"

func main() {
	/*
	古典问题：有一对兔子，从出生后第3个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，
	假如兔子都不死，问每个月的兔子总数为多少？12个月
	 */

	//方法一
	a := 1 //1月
	b := 1 //2月
	c := 0 //当前月
	for i := 3; i <= 12; i++ {
		c = a + b
		a = b
		b = c
	}
	fmt.Println(c)

	//方法二
	f1 := 1 //单月
	f2 := 1 //双月
	for i := 3; i <= 12; i += 2 {
		f1 += f2
		f2 += f1
	}
	fmt.Println(f2)
}
