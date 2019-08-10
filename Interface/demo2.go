

package main

import "fmt"

type Programmer interface {
	Greeting() string
}

type GoProrammer struct{}

func (g *GoProrammer) Greeting() string {
	return "fmt.Println(\"Hello World\")"
}

func main() {
	var p Programmer
	p = new(GoProrammer)
	fmt.Println(p.Greeting())
}
