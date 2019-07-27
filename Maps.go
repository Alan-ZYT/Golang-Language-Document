/*

	key在map里面是无序的，这个map是一个哈希map
	创建：map[string]string{]
	获取元素：m[key]
	Key不存在时，获得Value类型的初始值
	用两个值接收它 value,ok:=m[key]来判断是否存在key

	用delete函数删除一个key

	Map的遍历
	使用range遍历key，或者遍历key，value对
	不保证遍历顺序，如需顺序遍历，需手动对key排序

	使用len获取元素个数

	map的key
	map使用哈希表，key必须可以比较相等
	除了slice，map，function的内建类型都可以作为key
	struct类型不包含上述字段，也可以作为key

	练习题：寻找最长不含有重复字符的子串 https://leetcode.com
    Example1：abcabcbb --> abc
    Example1：bbbbb --> b
    Example1：pwwkew --> wke

*/

package main

import "fmt"

//练习题：寻找最长不含有重复字符的子串
func lengthOfNonRepeatingSUbStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
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

	fmt.Println(lengthOfNonRepeatingSUbStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSUbStr("a"))
	fmt.Println(lengthOfNonRepeatingSUbStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSUbStr(""))
	fmt.Println(lengthOfNonRepeatingSUbStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSUbStr("pwwkew"))

	m1 := map[string]string{
		"name":    "alan",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]string)
	var m3 map[string]string
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("-----------------")
	fmt.Println("Traversing map...")
	fmt.Println("-----------------")

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("-----------------")
	fmt.Println("Getting values...")
	fmt.Println("-----------------")

	if courseName, ok := m1["course"]; ok {
		fmt.Println(courseName, ok)
	} else {
		fmt.Println("key does not exist!!!")
	}

	fmt.Println("-----------------")
	fmt.Println("Delete values...")
	fmt.Println("-----------------")

	name, ok := m1["name"]
	fmt.Println(name, ok)
	delete(m1, "name")
	name, ok = m1["name"]
	fmt.Println(name, ok)

}
