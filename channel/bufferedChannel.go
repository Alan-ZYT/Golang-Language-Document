/*
缓冲 cahnnel：
	向无缓冲channel发送数据，运行程序会deadlock，因为没有接收者，一旦发数据，必须要有接收者；

思考：
	channel一旦发送数据，就一定要进行go程的切换比较耗费资源，虽然go程是轻量级的，但是发完1，必须等待接收，比较耗资源；
	可以加入一个缓冲区；
*/

package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d recevied %c\n", id, <-c)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func bufferedChanel() {
	ch := make(chan int, 3)
	go worker(0, ch)
	ch <- 'a'
	ch <- 'b'
	ch <- 'c'
	ch <- 'd'
	time.Sleep(time.Millisecond)
}

func main() {
	bufferedChanel()
}
