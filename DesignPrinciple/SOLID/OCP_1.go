package main

import "fmt"

type UG struct{}

func (b *UG) eat() {
	fmt.Println("Saving...")
}

func (b *UG) code() {
	fmt.Println("Transferring")
}

func (b *UG) sleep() {
	fmt.Println("Paying...")
}

func main() {
	ug := &UG{}

	ug.eat()
	ug.code()
	ug.sleep()

	fmt.Println("Why interface?")
}
