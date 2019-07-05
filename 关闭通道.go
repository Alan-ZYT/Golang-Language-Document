package main

/*
	关闭通道:
	从通道中读取数据：
		data := <- chan
		data,ok := <- chan
			同value,ok:=map[key]类似

		ok的数值为true，通道正常的，读取到的data数据有效可用
		ok的数值为false，通道关闭，读取到data是类型的零值。

	通道的关闭：发送方如果数据写入完毕，可以关闭通道，用于同时接收方数据传递完毕。
		ch1-->写入数据

		ch2-->读取数据

*/

import "fmt"

//关闭channel
func main() {

	ch5 := make(chan int)
	go sendData(ch5)

	//读取数据
	for {
		data, ok := <-ch5
		if !ok {
			fmt.Println("读取数据完毕,通道关闭了...")
			break
		}
		fmt.Println("main中读取到数据是:", data, ok)
	}
}

func sendData(ch5 chan int) {
	for i := 1; i <= 10; i++ {
		ch5 <- i
	}
	fmt.Println("发送方写入数据完毕...")
	close(ch5)
}
