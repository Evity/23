package main

import "fmt"

//A class should have one, and only one, reason to change.
//Single responsibility principle

/* type Clothes struct{}

func (c *Clothes) OnWork() {
	fmt.Println("Working Uniform")
}

func (c *Clothes) OnShop() {
	fmt.Println("Shopping Clothes")
}

func main() {
	c := Clothes{}

	// Working
	fmt.Println("Working...")
	c.OnWork()

	// Shopping
	fmt.Println("Shopping...")
	c.OnShop()
}
*/

type ClothesShop struct{}

func (cs *ClothesShop) Style() {
	fmt.Println("On shopping")
}

type ClothesWork struct{}

func (cw *ClothesWork) Style() {
	fmt.Println("On working")
}

func main() {
	//Working tx
	cw := ClothesWork{}
	cw.Style()

	//Shopping tx
	cs := ClothesShop{}
	cs.Style()
}
