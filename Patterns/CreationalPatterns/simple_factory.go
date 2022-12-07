package main

import "fmt"

type Ball struct {
	name string
	//...
}

//the simple factory violates SRP and OCP, but it is easy to understand and use.
//the simple factory compromise the availability of reuse and maintainability of classes.

func (b *Ball) Bounce(name string) {
	if name == "Basketball" {
		fmt.Println("Basketball bounce")

	} else if name == "Football" {
		fmt.Println("Football bounce")
	} else if name == "Tennis" {
		fmt.Println("Tennis bounce")
	} else if name == "Volleyball" {
		fmt.Println("Volleyball bounce")
	} else if name == "Badminton" {
		fmt.Println("Badminton bounce")
	} else {
		fmt.Println("Unknown ball")
	}
}

func NewBall(name string) *Ball {
	ball := new(Ball)

	if name == "Basketball" {
		ball.name = "Basketball"
	} else if name == "Football" {
		ball.name = "Football"
	} else if name == "Tennis" {
		ball.name = "Tennis"
	} else if name == "Volleyball" {
		ball.name = "Volleyball"
	} else if name == "Badminton" {
		ball.name = "Badminton"
	} else {
		ball.name = "Unknown"
	}

	return ball
}

func main() {
	basketball := NewBall("Basketball")
	basketball.Bounce("Basketball")

	football := NewBall("Football")
	football.Bounce("Football")

	tennis := NewBall("Tennis")
	tennis.Bounce("Tennis")

	volleyball := NewBall("Volleyball")
	volleyball.Bounce("Volleyball")

	badminton := NewBall("Badminton")
	badminton.Bounce("Badminton")

	unknown := NewBall("Unknown")
	unknown.Bounce("Unknown")
}
