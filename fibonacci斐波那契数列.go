package main

import "fmt"

func fibonacci(chr <-chan int, quit chan<- string) {

	for {
		select {
		case num := <-chr:
			fmt.Println("num=", num)
		case quit <- "exit":
			return
		}
	}
}

func main() {

	ch := make(chan int)
	quit := make(chan string)

	go fibonacci(ch, quit)

	x, y := 1, 1

	for i := 0; i < 10; i++ {
		ch <- x
		x, y = y, x+y
	}

	<-quit
}
