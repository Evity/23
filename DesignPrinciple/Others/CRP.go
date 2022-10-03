package main

import "fmt"

//Composite Reuse Principle
//Classes that are used together are packaged together.
//Composition is higher privilege than inherits

type Carrot struct{}

func (c *Carrot) Sing() {
	fmt.Println("Yo.Yo..")
}

type childCarrot struct {
	Carrot
}

func (cc *childCarrot) Imitate() {
	fmt.Println("Hello!")
}

type fusionCarrot struct {
	c *Carrot
}

func (fc *fusionCarrot) Imitate() { //composite is less coupling
	fmt.Println("Hello!")
}

func (fc *fusionCarrot) Sing(c *Carrot) { //least coupling
	c.Sing()
}

func main() {
	//carrot polymorph by inherit
	childCarrot := new(childCarrot)
	childCarrot.Sing()
	childCarrot.Imitate()

	//carrot polymorph by composite, the composite way is better choice
	fusionCarrot := new(fusionCarrot)
	fusionCarrot.c.Sing()
	fusionCarrot.Imitate()
	fusionCarrot.Sing(&Carrot{})
}
