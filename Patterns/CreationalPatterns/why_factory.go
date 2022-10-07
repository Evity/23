package main

import "fmt"

type Fruit struct {
}

func (f *Fruit) Show(name string) {
	//violates SRP, OCP
	if name == "apple" {
		fmt.Println("It's apple!")
	} else if name == "banana" {
		fmt.Println("It's banana!")
	} else {
		fmt.Println("It's general fruit!")
	}
}

func NewFruit(name string) *Fruit {
	//violate OCP, it modify existed code to add fruit
	fruit := new(Fruit)

	if name == "apple" {

	} else if name == "banana" {

	} else {

	}

	return fruit
}

func main() {
	//Too much coupling in this file!

	apple := NewFruit("apple")
	apple.Show("apple")

	banana := NewFruit("banana")
	banana.Show("banana")

	pear := NewFruit("pear")
	pear.Show("pear")
}
