package main

import "fmt"

func main() {
	/*
	作业11：		行	空格		×		换行
    *			1	4		1		1
   ***			2	3		3		1
  *****			3	2		5		1
 *******		4	1		7		1
*********		5	0		9		1
 *******		1	1	2	7		1
  *****			2	2	4	5		1
   ***			3	3	6	3		1
    *			4	4	8	1		1
	 */
	 n := 0
	 fmt.Println("请输入菱形的边长：")
	 fmt.Scanln(&n)
	//上三角
	for i := 1; i <= n; i++ {
		//1.打印空格
		for j := 0; j < n-i; j++ {
			/*
			i=1	j=0,1,2,3	j<4
			i=2	j=0,1,2		j<3
			i=3 j=0,1		j<2
			 */
			fmt.Print(" ")
		}
		//2.打印×
		for k := 0; k < i*2-1; k++ {
			/*
			i=1	k=0			k<1
			i=2	k=0,1,2		k<3
			i=3 k=0,1,2,3,4	k<5
			 */
			fmt.Printf("*")
		}
		//3.换行
		fmt.Println()
	}
	//下三角
	for i:=1;i<n;i++{
		//空格
		for j:=0;j<i;j++{
			fmt.Print(" ")
		}
		//星星
		for k:=0;k<2*n-1-2*i;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
