/*

	slice底层是一个数组，有三个变量，其中有一个指针，指向slice开头的那个元素;
	有len，说明slice的长度是多少，使用方括号取值的时候，只能取到len里面的值，下标大于等于len，都会报错：下标越界;
	还有一个值cap，代表了整个数组从指针开始到结束的整个长度;

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr[2:6]
	s2 = s1[3:5] // [s1[3], s1[4]]

	s1的值为[2 3 4 5],s2的值为[5 6]
	Slice可以向后扩展，不可以向前扩展
	s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)

	append 添加元素时如果超越cap，系统会重新分配更大的底层数组
	由于值传递的关系，必须接收append的返回值
	Exanple：s = append(s,val)

*/

package main

import "fmt"

// 切片是引用类型
func updateSlice(s []int) {
	s[0] = 100
}

func main1() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr =", arr)
	s1 = arr[2:6]
	s2 = s1[3:5] // [s1[3], s1[4]]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	// s4 and s5 no longer view arr.
	fmt.Println("arr =", arr)

	// Uncomment to run sliceOps demo.
	// If we see undefined: sliceOps
	// please try go run slices.go sliceops.go
	fmt.Println("Uncomment to see sliceOps demo")
	// sliceOps()
}

func printSlice(s []int) {
	fmt.Printf("%v,Len=%d,Cap=%d\n", s, len(s), cap(s))
}

// create slice
func main() {
	fmt.Println("Createing Slice...")
	var s []int
	for i := 0; i < 20; i++ {
		s = append(s, 2*i+1)
		printSlice(s)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice...")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice...")
	//删除下标为8的元素
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front...")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Poping from back...")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}
