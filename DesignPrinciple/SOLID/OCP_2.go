package main

import "fmt"

//You should be able to extend a classes behavior, without modifying it

type UG interface {
	Do()
}

type LazyDog struct{}

func (ld *LazyDog) Do() {
	fmt.Println("Do nothing.")
}

type Study struct{}

func (s *Study) Do() {
	fmt.Println("Study!")
}

type SportsMan struct{}

func (sm *SportsMan) Do() {
	fmt.Println("Working out!")
}

func main() {
	ld := &LazyDog{}
	ld.Do()

	s := &Study{}
	s.Do()

	sm := &SportsMan{}
	sm.Do()
}
