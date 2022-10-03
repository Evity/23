package main

import "fmt"

//You should be able to extend a classes behavior, without modifying it

type UG struct{}

func (b *UG) eat() {
	fmt.Println("eating...")
}

func (b *UG) code() {
	fmt.Println("coding")
}

func (b *UG) sleep() {
	fmt.Println("sleeping...")
}

func main() {
	ug := &UG{}

	ug.eat()
	ug.code()
	ug.sleep()

	fmt.Println("Why interface?")
}
