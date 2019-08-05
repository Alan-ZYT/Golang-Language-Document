/*
	关闭channel，永远是发送方关闭的，通知接收方没有新的数据要发了；
	关闭的channel还可以收到数据，收到的永远是chan具体类型的零值
	理论基础：CSP模型
	不要通过共享内存来通信，通过通信来共享内存。

关闭cahnnel方法：

	一：
		n, ok := <-ch // n == 具体的值 ；ok == 是否还有值
		if !ok {
			break
		}

	二：
		for n := range ch {}

*/

package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	for {
		/*	n, ok := <-ch
			if !ok {
				break
			}*/
		for n := range ch {
			fmt.Printf("Worker %d recevied %d\n", id, n)
		}
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

// Close channel
func closeChanel() {
	ch := make(chan int, 3)
	go worker(0, ch)
	ch <- 'a'
	ch <- 'b'
	ch <- 'c'
	ch <- 'd'
	close(ch)
	time.Sleep(time.Millisecond)
}

func main() {
	closeChanel()
}
