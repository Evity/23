package main

type Cook struct{}

type Order interface {
	Serve()
}

func (c *Cook) BBQChicken() {
	println("Made Chicken!")
}

func (c *Cook) BBQPork() {
	println("Made Pork!")
}

type Waitress struct {
	orders []Order
}

func (w *Waitress) Serve() {
	if w.orders == nil {
		return
	}
	for _, order := range w.orders {
		order.Serve()
	}
}

type OrderChicken struct {
	cook *Cook
}

func (o *OrderChicken) Serve() {
	o.cook.BBQChicken()
}

type OrderPork struct {
	cook *Cook
}

func (o *OrderPork) Serve() {
	o.cook.BBQPork()
}

func main() {
	cook := new(Cook)

	oc := OrderChicken{cook}
	op := OrderPork{cook}

	w := new(Waitress)
	w.orders = append(w.orders, &oc)
	w.orders = append(w.orders, &op)

	w.Serve()
}
