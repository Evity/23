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

type clothesShop struct{}

func (cs *clothesShop) Style() {
	fmt.Println("On shopping")
}

type clothesWork struct{}

func (cw *clothesWork) Style() {
	fmt.Println("On working")
}

func main() {
	//Working tx
	cw := clothesWork{}
	cw.Style()

	//Shopping tx
	cs := clothesShop{}
	cs.Style()
}
