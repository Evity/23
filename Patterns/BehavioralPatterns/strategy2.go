package main

import "fmt"

type SellStrategy interface {
	GetPrice(price float64) float64
}

type NormalStrategy struct {
}

type DiscountStrategy struct {
}

func (ns *NormalStrategy) GetPrice(price float64) float64 {
	fmt.Println("NormalStrategy")
	return price
}

func (ds *DiscountStrategy) GetPrice(price float64) float64 {
	fmt.Println("DiscountStrategy")
	return price * 0.8
}

//============environment class================
type Goods struct {
	Price float64

	Strategy SellStrategy
}

func (g *Goods) SetStrategy(s SellStrategy) {
	g.Strategy = s
}

func (g *Goods) Sell() (price float64) {
	price = g.Strategy.GetPrice(g.Price)
	fmt.Println("Origin price: ", g.Price, "Discount price: ", price)
	return
}

func main() {
	nike := Goods{
		Price: 100,
	}

	nike.SetStrategy(new(NormalStrategy))
	nike.Sell()

	nike.SetStrategy(new(DiscountStrategy))
	nike.Sell()

}
