package main

import (
	"fmt"
)

func main() {
	/*
		a：先输出提示语句，并接受用户输入的年、月。
		b：根据用户输入的年，先判断是否是闰年。
		C：根据用户输入的月来判断月的天数。
		D：用循环计算用户输入的年份距1900年1月1日的总天数。
		E：用循环计算用户输入的月份距输入的年份的1月1日共有多少天。
		F：相加D与E的天数，得到总天数。
		G：用总天数来计算输入月的第一天的星期数。
		H：根据G的值，格式化输出这个月的日历！

	*/

	fmt.Println("请输入年份：")
	year := 2018
	fmt.Scanln(&year)
	fmt.Println("请输入月份：")
	month := 5
	fmt.Scanln(&month)
	//step1:求1900年距离year年month上一个月底的总天数
	days := 0 //总天数
	//整年
	for i := 1900; i < year; i++ {
		if i%4 == 0 && i%100 != 0 || i%400 == 0 { //闰年
			days += 366
		} else {
			days += 365
		}
	}
	//整月
	for i := 1; i < month; i++ {
		day := 0 //第i个月的天数
		switch i {
		case 1, 3, 5, 7, 8, 10, 12:
			day = 31
		case 4, 6, 9, 11:
			day = 30
		case 2:
			if year%4 == 0 && year%100 != 0 || year%400 == 0 {
				day = 29
			} else {
				day = 28
			}
		}
		days += day
	}
	//step2：求空格
	//kong := days % 7
	kong := (days + 1) % 7
	//step3：求当月的天数，打印输出即可
	day := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			day = 29
		} else {
			day = 28
		}
	}

	//fmt.Println("一\t二\t三\t四\t五\t六\t日")
	fmt.Println("日\t一\t二\t三\t四\t五\t六")
	for i := 0; i < kong; i++ {
		fmt.Print("\t")
	}
	for i := 1; i <= day; i++ {
		fmt.Print(i, "\t")
		if (i+kong)%7 == 0 {
			fmt.Println()
		}
	}

}
