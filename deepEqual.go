package main

import (
	"fmt"
	"reflect"
)

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func DeepEqualMap() {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}

	// fmt.Println(a == b) // invalid operation: a == b (map can only be compared to nil)
	fmt.Println("a==b?", reflect.DeepEqual(a, b))
}

func DeepEqualSLice() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	//fmt.Println("s1 == s2?", s1 == s2) // invalid operation: s1 == s2 (slice can only be compared to nil)
	//fmt.Println("s1 == s3?", s1 == s3) // invalid operation: s1 == s3 (slice can only be compared to nil)

	fmt.Println("s1 == s2?", reflect.DeepEqual(s1, s2))
	fmt.Println("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

func main() {
	DeepEqualMap()
	DeepEqualSLice()
}
