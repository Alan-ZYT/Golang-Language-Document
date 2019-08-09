// map排序：
// 先获取所有key，把key进行排序，再按照排序好的key，进行遍历。
package main

import (
	"fmt"
	"sort"
)

func main() {
	
	// 方法1
	m := map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "5"}

	var slice []string
	for k := range m {
		slice = append(slice, k)
	}
	fmt.Printf("slise string is : %v\n", slice)
	sort.Strings(slice[:])
	fmt.Printf("sorted slice string is : %v\n", slice)

	for _, v := range slice {
		fmt.Println(m[v])
	}
	
	// 方法2
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m1))
	i := 0
	for k := range m1 {
		s[i] = k
		i++
	}
	sort.Ints(s)
	for _, v := range s {
		fmt.Println(m1[v])
	}
}
