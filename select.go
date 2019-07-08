package main

/*
	select ：监听 channel 上的 数据流动（r、w）

	语法：
		select {
		case  channnel IO 操作:

		case  channnel IO 操作:
		.....
		default :
	        }

	执行流程：
		1.如果有多个 case 都可以执行，select 会随机执行一个可运行的 case。
		2.如果没有 case 可运行，如果有 default，执行 default，如果没有 default，它将阻塞，直到有 case 可运行。


select 使用注意事项：

	1. case分支中必须是一个 IO操作：

	2. 当case分支不满足监听条件，阻塞当前 case 分支

	3. 如果同时有多个case分支满足，select随机选定一个执行（ select底层实现，case对应一个 goroutine）

	4. 一次select 监听，只能执行一个case分支。通常将select放于for 循环中

	5. default 在所有case均不满足时，默认执行的分组，为防止忙轮询，通常将 for 里 select中的 default 省略。

	【重要结论】：使用 select的 go程，与其他go程间 采用  【异步通信】 通信方式。
	
	
select 实现的 超时：

	1. select {
	         case ch<-:
	         case <-time.After(时间) ：
	    } 

	2.  当 select 监听的 ch 写事件满足时。会重置 time.After 的定时器时间。

	3.  只有 select 监听的所有case均不满足，time.After 计时满 ，该 case 才满足监听条件（读事件）。

*/

import (
	"fmt"
	"time"
)

func main01() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()

	select {
	case data := <-ch1:
		fmt.Println("ch1中读取到数据:", data)
	case data := <-ch2:
		fmt.Println("ch2中读取到数据：", data)
	default:
		fmt.Println("执行了default...")
	}
}

func main02() {

	var ch = make(chan interface{})
	var quit = make(chan bool)

	go func() {

		for {
			select {
			case num := <-ch:
				fmt.Println("num=", num)
			case <-quit:
				//runtime.Goexit()
				goto AAA
			}
		}
	AAA:
		fmt.Println("子 goroutine...over...")
	}()

	for i := 0; i < 50; i++ {
		ch <- i
	}

	quit <- true

	time.Sleep(time.Second)
	fmt.Println("main...over...")
}

func main03() {
	var ch1 = make(chan interface{})
	var ch2 = make(chan interface{})

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i
		}
	}()

	select {
	case data1 := <-ch1:
		fmt.Println("read ch1 data", data1)
	case data2 := <-ch2:
		fmt.Println("read ch2 data", data2)
		//default
		//	fmt.Println("default...")
	}

	fmt.Println("mian...over...")
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 100
		//close(ch1)
	}()

	go func() {
		ch2 <- 200
	}()

	//结束循环的条件：通道结束，或者超时5次
	count := 0
out:
	for { //1次，2次，3次
		//time.Sleep(1*time.Second)
		select {
		case data, ok := <-ch1:
			if !ok {
				fmt.Println("通道关闭了...")
				break out
			}
			fmt.Println("ch1中读取数据：", data)
		case data := <-ch2:
			fmt.Println("ch2中读取数据：", data)
		case <-time.After(2 * time.Second):
			fmt.Println("超时2s...")
			count++
			if count == 5 {
				break out
			}
			//default:
			//	fmt.Println("default。。。")
		}
	}
}
