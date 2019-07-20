/*

	Go 中的 switch 做类型判断。

*/

package main

import "fmt"

func Istype(v interface{}) {
	switch v.(type) {
	case bool:
		fmt.Println("bool")
	case float64:
		fmt.Println("float64")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("default")
	}
}

func main() {
	Istype(1)      //int
	Istype("alan") //string
	Istype('q')    //default
	Istype(3.14)   //float64
	Istype(true)   //bool
	Istype(nil)    //nil
	i := 1
	Istype(&i)      //default
	fmt.Println(&i) //0xc0000200e0
}
