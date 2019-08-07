package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	/*
	剪刀石头布
	设计一款游戏：
	剪刀，石头，布，猜拳游戏
系统：剪刀，石头，布
玩家：剪刀，石头，布
1：剪刀
2：石头
3：布
	 */
	//step1,系统产生随机数，表示剪刀石头布
	rand.Seed(time.Now().UnixNano())
	choose := ""
	for {
		fmt.Println("n代表退出，任意键代表游戏：")
		fmt.Scanln(&choose)
		if choose == "n" {
			fmt.Println("您选择退出游戏，欢迎下次光临。。")
			break
		} else {
			//游戏
			sysNum := rand.Intn(3) + 1

			//step2，玩家输入：剪刀，石头，布
			num := 0
			fmt.Println("请输入：1代表剪刀，2代表石头，3代表布：")
			fmt.Scanln(&num)

			//step3，系统玩家
			name1 := ""
			name2 := ""
			switch sysNum {
			case 1:
				name1 = "剪刀"
			case 2:
				name1 = "石头"
			case 3:
				name1 = "布"
			}
			switch num {
			case 1:
				name2 = "剪刀"
			case 2:
				name2 = "石头"
			case 3:
				name2 = "布"
			}
			fmt.Printf("系统出拳：%s，玩家出拳：%s\n", name1, name2)

			//step4：判断输赢
			if sysNum == num {
				fmt.Println("\t平局。。。")
			} else {
				switch num {
				case 1: //玩家，剪刀
					if sysNum == 2 { //系统，石头
						fmt.Println("\t玩家输了，呜呜。。。。")
					} else { //系统，布
						//sysNum=3
						fmt.Println("\t玩家赢了，哈哈。。。")
					}
				case 2: //玩家，石头
					if sysNum == 1 { //系统，剪刀
						fmt.Println("\t玩家赢了。。")

					} else {
						fmt.Println("玩家输了。。")
					}
				case 3:
					if sysNum == 1 {
						fmt.Println("\t玩家输了。。")
					} else {
						fmt.Println("\t玩家赢了。。。")
					}
				}
			}
		}
	}

}
