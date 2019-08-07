package main

import "fmt"

func main() {
	/*
		练习4：给定一个数组，arr1 := [10] int{5,4,3,7,10,2,9,8,6,1}，求数组中所有数据的总和。
	*/
	arr1 := [10]int{5, 4, 3, 7, 10, 2, 9, 8, 6, 1}

	sum := 0

	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	fmt.Println(sum)
}
