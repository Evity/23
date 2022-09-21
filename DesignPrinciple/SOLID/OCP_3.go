package main

import "fmt"

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

func DoThings(u UG) {
	u.Do()
}

func main() {
	DoThings(&LazyDog{})
	DoThings(&Study{})
	DoThings(&SportsMan{})
}
