package main

import "fmt"

//A reverse example of DIP

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

func (s *Soldier) EquipAR(ar *AR) {
	fmt.Println("Soldier equip AR-15")
	ar.Fire()
}

func (s *Soldier) EquipAK(ak *AK) {
	fmt.Println("Soldier equip AR-15")
	ak.Fire()
}

type General struct {
}

func (g *General) EquipAR(ar *AR) {
	fmt.Println("Soldier equip AR-15")
	ar.Fire()
}

func (g *General) EquipAK(ak *AK) {
	fmt.Println("Soldier equip AR-15")
	ak.Fire()
}

func main() {
	ak := &AK{}
	rookie := &Soldier{}
	rookie.EquipAK(ak)

	ar := &AR{}
	general := &General{}
	general.EquipAR(ar)
}
