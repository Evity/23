package main

import "fmt"

//Classes that are used together are packaged together.

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
	c Carrot
}

func (fc *fusionCarrot) Imitate() {
	fmt.Println("Hello!")
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
}
