package main

type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()

	AddHook() bool //hook
}

type template struct {
	b Beverage
}

func (t *template) PrepareRecipe() {
	if t == nil || t.b == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	if t.b.AddHook() {
		t.b.AddCondiments()
	}
}

type Coffee struct {
	template
}

func NewCoffee() *Coffee {
	return &Coffee{template{&Coffee{}}}
}

func (c *Coffee) BoilWater() {
	println("Boiling water")
}

func (c *Coffee) Brew() {
	println("Dripping Coffee through filter")
}

func (c *Coffee) PourInCup() {
	println("Pouring into cup")
}

func (c *Coffee) AddCondiments() {
	println("Adding Sugar and Milk")
}

func (c *Coffee) AddHook() bool {
	return true
}

type Tea struct {
	template
}

func NewTea() *Tea {
	return &Tea{template{&Tea{}}}
}

func (t *Tea) BoilWater() {
	println("Boiling water")
}

func (t *Tea) Brew() {
	println("Steeping the tea")
}

func (t *Tea) PourInCup() {
	println("Pouring into cup")
}

func (t *Tea) AddCondiments() {
	println("Adding Lemon")
}

func (t *Tea) AddHook() bool {
	return false
}

func main() {
	coffee := NewCoffee()
	coffee.PrepareRecipe()

	tea := NewTea()
	tea.PrepareRecipe()
}
