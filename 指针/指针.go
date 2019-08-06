/*
 Go语言的指针不能参与运算

 Go语言的参数传递是值传递还是引用传递？？？
 C,C++既可以值传递又可以引用传递;
 Java,Python绝大部分类型是引用传递;

 Golang只有只有值传递一种方式

 值传递？ 性能下降
 相当于做了一份拷贝，不会改变原来的数据

 引用传递？ 利用指针
 不做拷贝，指向同一个变量(内存地址)，会改变原来的数据

*/

package main

import (
	"fmt"
)

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 300
	fmt.Println(a)

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	fmt.Println(swap2(3, 4))
}

//交换两个变量的值 方法1
func swap(a, b *int) {
	*b, *a = *a, *b
}

//方法2
func swap2(a, b int) (int, int) {
	return b, a
}
