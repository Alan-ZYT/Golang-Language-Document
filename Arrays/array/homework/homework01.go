package main

import "fmt"

func main() {
	/*
		作业4：操场上有一群人，人数在100到200之间。三人一组多1人，四人一组多2人，五人一组多3人。问操场上共有多少人。

		思路：
			变量：i
				100,200
	*/
	for i := 100; i <= 200; i++ {
		if i%3 == 1 && i%4 == 2 && i%5 == 3 {
			fmt.Println(i)
		}

	}
}