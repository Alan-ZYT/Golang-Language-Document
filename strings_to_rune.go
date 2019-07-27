/*

	使用range遍历pos，rune对
	使用utf8.RuneCountInString获得字符数量
	使用len获得字节长度
	使用[]byte获得字节

*/
package main

import (
	"fmt"
	"unicode/utf8"
)

//练习题：寻找最长不含有重复字符的子串，使用rune类型  -->国际版
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println("寻找最长不含有重复字符的子串，使用rune类型  -->国际版")
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("a"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(lengthOfNonRepeatingSubStr("中国欢迎你"))
	fmt.Println(lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
	fmt.Println()

	//底层处理方法
	str := "中国欢迎你!"
	for _, b := range []byte(str) {
		fmt.Printf("%X ", b)
		// E4 B8 AD E5 9B BD E6 AC A2 E8 BF 8E E4 BD A0 21  -->UTF8编码，可变长
	}
	fmt.Println()
	for i, ch := range str { //ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
		// (0 4E2D) (3 56FD) (6 6B22) (9 8FCE) (12 4F60) (15 21)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(str)) // Rune count: 6

	bytes := []byte(str)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	//上层处理方法
	for i, ch := range []rune(str) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}
