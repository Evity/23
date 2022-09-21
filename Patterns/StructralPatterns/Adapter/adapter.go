package main

import "fmt"

// Object to adapt
type BigBrother interface {
	Watch()
}

// tx class, depend on bigBrother
type Governor struct {
	boss BigBrother
}

func NewGovernor(b BigBrother) *Governor {
	return &Governor{b}
}

func (g *Governor) Govern() {
	fmt.Println("Governing")
	g.boss.Watch()
}

// object be adapted
type People struct{}

func (p *People) Live() {
	fmt.Println("People surviving")
}

// adapter
type Monitor struct {
	p *People
}

func (m *Monitor) Watch() {
	fmt.Println("Big brother is watching you!")

	//call adapter's method
	m.p.Live()
}

func NewMonitor(p *People) *Monitor {
	return &Monitor{p}
}

// tx logic layer
func main() {
	party := NewGovernor(NewMonitor(new(People)))

	party.Govern()
}
