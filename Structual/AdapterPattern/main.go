package main

import "fmt"

type Mobile interface {
	chargeMobile()
}

type apple struct {}

func (a *apple) chargeMobile() {
	fmt.Println("Apple mobile is charging...")
}
// Adaptee
type samsung struct {}

func (s *samsung) chargeMobile() {
	fmt.Println("Samsung mobile is charging...")
}
// samsung adapter
type samsungAdapter struct {
	samsungMobile *samsung
}

func (s *samsungAdapter) chargeMobile() {
	s.samsungMobile.chargeMobile()
}

// client
type client struct {}

func (c *client) chargeMobile(mob Mobile) {
	mob.chargeMobile()
}

func main() {
	// First
	apple := &apple{}
	client := &client{}
	client.chargeMobile(apple)

	samsung := &samsung{}
	samsungAdapter := &samsungAdapter{samsungMobile: samsung}
	client.chargeMobile(samsungAdapter)
}

