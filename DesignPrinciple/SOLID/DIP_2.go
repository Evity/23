package main

import "fmt"

/*
Abstract the system into 3 layers

tx logic layer on top
 ->final concrete implementation

abstract layer in the middle
 ->first layer to construct

impl layer at the bottom
 ->impl how abstract layer actually work
*/

// ab layer
type militaryRank interface {
	Equip(ar assaultRifle)
}

type assaultRifle interface {
	Fire()
}

// impl layer
type AR struct {
}

func (a *AR) Fire() {
	fmt.Println("AR-15 firing...")
}

type AK struct {
}

func (a *AK) Fire() {
	fmt.Println("AKM firing...")
}

type Soldier struct {
}

type General struct {
}

func (s *Soldier) Equip(ar assaultRifle) {
	fmt.Println("Soldier equip gun")
	ar.Fire()
}

func (g *General) Equip(ar assaultRifle) {
	fmt.Println("Soldier equip gun")
	ar.Fire()
}

// tx layer
func main() {
	var ar15 assaultRifle
	ar15 = &AR{}

	var rookie militaryRank
	rookie = &Soldier{}

	rookie.Equip(ar15)

	var ak assaultRifle
	ak = &AK{}

	var general militaryRank
	general = &General{}

	general.Equip(ak)
}
