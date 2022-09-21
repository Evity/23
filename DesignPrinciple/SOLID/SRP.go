package main

import "fmt"

type Cat struct{}

func (c *Cat) Roar() {
	fmt.Println("Meow~")
}

type Carrot struct{}

func (c *Carrot) Imitate() {
	fmt.Println("Hello!")
}

func main() {
	//carrot polymorph
	carrot := new(Carrot)
	carrot.Imitate()

	//cat polymorph
	cat := new(Cat)
	cat.Roar()
}

//A class should have one, and only one, reason to change.
