package main

import "fmt"

func main() {
	/*
		递归函数(recursion)：一个函数自己调用自己，就叫做递归函数。
			递归函数要有一个出口，逐渐的向出口靠近,否则就进入死递归(死循环)
	*/

	/*
					求1-5的和:15    1 + 2 + 3 + 4 + 5 = 15
										getSum(5)
											getSum(4) + 5 相当于: 10 + 5

					求1-4的和:10  1 + 2 + 3 + 4 = 10
										getSum(4)
										getSum(3) + 4 相当于: 6 + 4 = 10

					求1-3的和:6  1 + 2 + 3 = 6
										getSum(3)
										getSum(2) + 3 相当于: 3 + 3 = 6


				   求1-2的和:3  1 + 2 = 3
										getSum(2)
						       			getSum(1) + 2 相当于:  1 + 2 = 3

				   getSum(1)的和: 1

		        	得到一个递归求和公式: getSum(5) = getSum(4) + 5
					如果把5看作是变量n:  getSum(n) = getSum(n-1) + n
	*/

	//1.求1-5的和
	sum := getSum(5)
	fmt.Println("1-5的和:", sum) //1-5的和: 15

	//2.fibonacci数列 第三位是前两位数的和:
	/*
		1	2	3	4	5	6	7	8	9	10	11	12		。。。
		1	1	2	3	5	8	13	21	34	55	89	144
	*/

	res := getFibonacci(12)
	fmt.Println("第12位数字:", res) //第12位数字: 144
}

func getFibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return getFibonacci(n-1) + getFibonacci(n-2)
}

func getSum(n int) int {
	fmt.Println("getSum函数执行了五次**********")
	if n == 1 {
		return 1
	}
	return getSum(n-1) + n
}
