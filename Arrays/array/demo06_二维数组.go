package main

import "fmt"

func main() {
	/*
		二维数组：存储的是一维的一维
			arr :=[4][5]int
			该二维数组的长度，就是4
			存储的元素是一维数组，一维数组的元素是数值
	*/

	arr11 := [4][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}
	fmt.Println(arr11)

	fmt.Println("二维数组的长度：", len(arr11))    //一维的个数 4
	fmt.Println("一维数组的长度：", len(arr11[0])) //一维中数值的个数 5

	fmt.Println(arr11[0][0]) //1
	fmt.Println(arr11[2][2]) //13
	fmt.Println(arr11[2][1]) //12
	fmt.Println(arr11[1][3]) //9

	//遍历二维数组
	for i := 0; i < len(arr11); i++ {
		for j := 0; j < len(arr11[i]); j++ {
			fmt.Print(arr11[i][j], "\t")
		}
		fmt.Println()
	}

	//for range遍历二维数组
	for _, arr111 := range arr11 {
		//fmt.Println(arr111)
		for _, val := range arr111 {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}
}
