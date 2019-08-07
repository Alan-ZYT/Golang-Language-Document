package main

import "fmt"

func main() {
	/*
		数组的遍历：
			依次访问数组中的元素
			方法一：arr[0],arr[1],arr[2]....

			方法二：通过循环，配合下标
				for i:=0;i<len(arr);i++{
					arr[i]
				}

			方法三：使用range
				range，词义"范围"
				不需要操作数组的下标，到达数组的末尾，自动结束for range循环。
					每次都数组中获取下标和对应的数值。
	*/

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range arr {
		fmt.Println("下标是：", index, ",数值是：", value)
	}

	arr1 := [10]int{5, 4, 3, 7, 10, 2, 9, 8, 6, 1}
	sum := 0
	//可以使用_，来舍弃下标
	for _, v := range arr1 {
		//fmt.Println(v)
		sum += v
	}
	fmt.Println(sum)
}
