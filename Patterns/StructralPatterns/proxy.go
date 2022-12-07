package main

import "fmt"

type Goods struct {
	Kind           string
	Authentication bool
}

type Shopping interface {
	Buy(goods *Goods)
}

type KoreaShopping struct{}

func (k *KoreaShopping) Buy(goods *Goods) {
	println("Buy goods in Korea")
}

type AmericaShopping struct{}

func (a *AmericaShopping) Buy(goods *Goods) {
	println("Buy goods in America")
}

type Proxy struct {
	shopping Shopping
}

func (p *Proxy) Buy(goods *Goods) {
	if p.distinguish(goods) {
		p.shopping.Buy(goods)

		p.check(goods)
	} else {
		println("Authentication failed")
	}
}

func (p *Proxy) distinguish(g *Goods) bool {
	fmt.Println("Distinguish goods: ", g.Kind)
	if !g.Authentication {
		fmt.Println("Fake!")
	}
	return g.Authentication
}

func (p *Proxy) check(g *Goods) {
	fmt.Println("Check goods: ", g.Kind)
}

func main() {
	goods1 := Goods{
		Kind:           "Apple",
		Authentication: true,
	}

	goods2 := Goods{
		Kind:           "HUAWEI 14 PRO MAX",
		Authentication: false,
	}

	//if not using proxy, need to specify the shopping type
	var shopping Shopping = new(KoreaShopping)

	if goods1.Authentication {
		fmt.Println("Authentication success: ", goods1.Kind)

		shopping.Buy(&goods1)

		fmt.Println("Check goods: ", goods1.Kind)
	}

	fmt.Println("Using Proxy")

	var proxy Shopping = new(Proxy)

	proxy.Buy(&goods1)
	proxy.Buy(&goods2)
}
