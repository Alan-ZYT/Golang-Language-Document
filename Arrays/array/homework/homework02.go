package main

import "fmt"

func main() {
	/*
		两个自然数x，y相除，商3余10，被除数，除数，商，余数的和是163，求被除数，除数

		"暴力破解法"，依次尝试可能的结果
	*/

	for x := 1; x < 150; x++ {
		y := 150 - x
		if x/y == 3 && x%y == 10 {
			fmt.Printf("被除数：%d，除数：%d\n", x, y)
		}
	}
}
