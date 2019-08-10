/*
	Go语言中，不需要显式声明实现了哪一个接口，只要有了这个接口定义的方法，就默认实现了这个接口，不需要告诉编译器;
	Go语言 Duck Type（鸭子类型）式接⼝实现;长得像鸭子，羽毛，扁嘴
	Go语言中所有的类型都实现了空接口
*/

package main

import "fmt"

type Empty interface{}

type USB interface {
	Name() string
	Connect()
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connecter:", pc.name)
}

func Disconnecter(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnectd:", v.name)
	default:
		fmt.Println("Unknown decive.")
	}
}

func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnecter(a)
}
