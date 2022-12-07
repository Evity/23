package main

import "fmt"

type V5 interface {
	Charge()
}

type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("Phone Charge")
	p.v.Charge()
}

type V220 struct{}

func (v *V220) Charge() {
	fmt.Println("V220 Charge")
}

type Adapter struct {
	v220 *V220
}

func (a *Adapter) Charge() {
	fmt.Println("Adapter Charge")
	a.v220.Charge()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

func main() {
	iphone := NewPhone(NewAdapter(&V220{}))

	iphone.Charge()
}
