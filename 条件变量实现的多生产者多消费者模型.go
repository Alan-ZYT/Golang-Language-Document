package main

/*
条件变量：
	Cond 语法：

	type Cond struct {
   		noCopy noCopy

   		L Locker		// 锁：互斥、读写

   		notify  notifyList
   		checker copyChecker
	}

	Cond.Wait():
		1)  阻塞等待条件变量满足

		2）释放已经掌握的互斥锁  （1/2 两步为 原子操作。）

			。。。 阻塞等。

		3） 被唤醒时，解除阻塞，重新加锁

	Cond.Signal()：

		唤醒阻塞在条件变量上的 一个 对象。

	Cond.Broadcast()

		唤醒阻塞在条件变量上的 所有 对象。

条件变量使用流程： --- 生产者为例

	1. 创建条件变量	var cond Sync.Cond	—— 结构体

	2. 初始化 条件变量使用的 锁 	cond.L = new(sysc.Mutex)

	3. 给条件变量的 互斥锁 加锁	cond.L.Lock()

	4. 判断条件变量是否满足。 调用wait ：1） 2） 3）

		for len(ch) == cap(ch) {   		// 此处判断，必须使用 for ，而不能使用 if
			cond.Wait()
		}

	[重要结论] 所有条件变量中关于 cond.Wait 的判断语句采用for语句去做.


	5. 生产数据：	产生随机数 num := rand.Intn()

	6. 写入公共区（缓冲区）	ch <- num

	7. 唤醒对端（消费者）	cond.Signal()

	8. 给条件变量的 互斥锁 解锁	cond.L.UnLock()

伪代码展示:

	c.L.Lock()
	for !condition() {
		c.Wait()
	}
	... make use of condition ...
	c.L.Unlock()

*/

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//创建一个全局的条件变量
var cond sync.Cond

//生产者
func producer(in chan<- interface{}, idx interface{}) {

	for {
		//给条件变量的成员 L 加锁
		cond.L.Lock()
		//判断条件变量是否满足
		for len(in) == cap(in) {
			cond.Wait()
		}
		//产生数据 -- 这里使用随机数
		data := rand.Intn(1000)
		//写入公共区
		in <- data
		fmt.Printf("%d写出: %d\n", idx, data)
		//发送信号唤醒阻塞在条件变量上的消费者
		cond.Signal()
		//解锁
		cond.L.Unlock()
	}
}

//消费者
func consumer(out <-chan interface{}, idx interface{}) {

	for {
		cond.L.Lock()

		for len(out) == 0 {
			cond.Wait()
		}

		//消费数据
		fmt.Printf("---%d---读到: %d\n", idx, <-out)
		//发送信号
		cond.Signal()
		cond.L.Unlock()
	}
}

func main() {
	//随机数方法
	rand.Seed(time.Now().UnixNano())
	//给条件变量的成员 L 初始化互斥锁
	cond.L = new(sync.Mutex)
	//有缓冲的公共区
	ch := make(chan interface{}, 10)

	for i := 0; i < 10; i++ {
		go producer(ch, i+1)
	}

	for i := 0; i < 10; i++ {
		go consumer(ch, i+1)
	}

	for {
		runtime.GC()
	}
}
