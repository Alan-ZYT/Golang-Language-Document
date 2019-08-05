/*
	channel是goroutine与goroutine之间的交互，必须使用另外一个goroutine接收它，发送一个数据必须要有接收者；
	否则 fatal error: all goroutines are asleep - deadlock!

	Go语言函数是一等功名，函数能够作为参数，能够作为返回值；
	channel也是一等功名，也能作为返回值，也能作为返回值
*/

package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("worker %d received %c\n", id, <-c)
	}
}

func chanDemo() {
	//var c chan int // 只是定义了变量c的类型是 chan int；c == nil
	//ch := make(chan int)
	//ch <- 1   //发送数据
	//n := <-ch // 接收数据

	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
