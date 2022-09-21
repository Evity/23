package main

import "fmt"

/*
Each unit should have only limited knowledge about other units:
 only units "closely" related to the current unit. (Black box)
*/

type Dev struct {
}

type Client struct {
}

type PM struct {
}

func (d *Dev) Code() {
	fmt.Println("Dev got request. Coding ...")
}

func (p *PM) PostRequest() {
	fmt.Println("PM post a request")
}

func (c *Client) Order() {
	fmt.Println("Client made an order")
}

func OrderRoutine() {
	dev := &Dev{}
	c := &Client{}
	pm := *&PM{}
	c.Order()
	pm.PostRequest()
	dev.Code()
	fmt.Println("Routine finished")
}

func main() {
	OrderRoutine()
}
