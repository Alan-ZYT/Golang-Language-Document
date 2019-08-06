package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"Unsupported operation: %s", op)
	}
}

// 15 / 7 = 2 ... 1
func div(a, b int) (q, r int) {
	// return a / b, a % b
	q = a / b
	r = a % b
	return
}

// 函数式编程 函数的参数是一个函数
func apply(op func(int, int) int, a, b int) int { //op是一个函数，有两个参数，返回int，接收两个参数a，b，返回int
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

// 可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

// 交换两个变量的值
func swap(a, b int) (int, int) {
	return b, a
}

func main() {

	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error handling:", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(15, 7)
	fmt.Printf("quotient: %d\t remainder: %d\n", q, r)

	fmt.Println("pow(3, 4) is:", apply(
		func(a int, b int) int {
			return int(math.Pow( //math.Pow  pow() 方法可返回 x 的 y 次幂的值
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1 + 2 + ... + 5 =", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	//a, b = b, a
	a, b = swap(a, b)
	fmt.Println("a, b after swap is:", a, b)
}
