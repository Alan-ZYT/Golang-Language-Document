//map排序：
//先获取所有key，把key进行排序，再按照排序好的key，进行遍历。
package main

import (
	"fmt"
	"sort"
)

func main() {
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
}
