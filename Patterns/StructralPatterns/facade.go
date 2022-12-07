package main

import "fmt"

type SubSystemA struct{}

func (sa *SubSystemA) MethodA() {
	println("SubSystemA Do")
}

type SubSystemB struct{}

func (sb *SubSystemB) MethodB() {
	println("SubSystemB Do")
}

type SubSystemC struct{}

func (sc *SubSystemC) MethodC() {
	println("SubSystemC Do")
}

type SubSystemD struct{}

func (sd *SubSystemD) MethodD() {
	println("SubSystemD Do")
}

type Facade struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
	d *SubSystemD
}

func (f *Facade) Method1() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f *Facade) Method2() {
	f.c.MethodC()
	f.d.MethodD()
}

func main() {
	sa := new(SubSystemA)
	sa.MethodA()
	sb := new(SubSystemB)
	sb.MethodB()

	fmt.Println("------------")

	f := Facade{
		a: new(SubSystemA),
		b: new(SubSystemB),
		c: new(SubSystemC),
		d: new(SubSystemD),
	}

	f.Method1()
}
