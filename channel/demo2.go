/*
	channel 是goroutine与goroutine之间的交互，必须使用另外一个goroutine接收它
	Go语言函数是一等功名，函数能够作为参数，能够作为返回值；
	channel也是一等功名，也能作为返回值，也能作为返回值

	下面的demo：
	channel 作为返回值

*/

package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo2() {
	//var c chan int // c == nil
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i // send data
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo2()
}
