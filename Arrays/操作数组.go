package main

import "fmt"

// 求出数组中最大Value
func main() {

	numbers := [6]int{2, 10, 8, 12, 6, 4}
	maxI := -1
	maxValue := -1

	for i, v := range numbers { // rang 关键字获取数组下标和value
		if v > maxValue {
			maxI, maxValue = i, v
		}
	}
	fmt.Printf("numbers Arrays maxIndex: %d\tmaxValue: %d\n", maxI, maxValue)

	// 数组求和
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println("numbers Arrays summation:", sum)
}
