package main

import "fmt"

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

func main() {
	//var dog Dog
	//dog.Pet.SpeakTo("Zhang")
	dog := new(Dog)
	dog.SpeakTo("Zhang")
}
