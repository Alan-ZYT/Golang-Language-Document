package main

import "fmt"

type Code string

type Programming interface {
	WriteGreeting() Code
}

type GoProgramming struct{}
type JavaProgramming struct{}

func (g *GoProgramming) WriteGreeting() Code {
	return "fmt.Println(\"Hello World!\")"
}

func (j *JavaProgramming) WriteGreeting() Code {
	return "System.out.Println(\"Hello World!\")"
}

func WriteFirstProgrammer(p Programming) {
	fmt.Printf("%T %v\n", p, p.WriteGreeting())
}

func main() {
	Gopher := &GoProgramming{}
	Javapher := new(JavaProgramming)

	WriteFirstProgrammer(Gopher)
	WriteFirstProgrammer(Javapher)

}

/*
*main.GoProgramming fmt.Println("Hello World!")
*main.JavaProgramming System.out.Println("Hello World!")
 */
