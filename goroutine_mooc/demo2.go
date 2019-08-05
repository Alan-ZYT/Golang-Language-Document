package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++            //交不出控制权
				runtime.Gosched() //演示效果，交出控制权
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
