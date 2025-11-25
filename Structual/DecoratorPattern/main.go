package main

import "fmt"

type Coffee interface {
	GetCost() float64
	GetDescription() string
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) GetCost() float64 {
	return 10.0
}

func (c *SimpleCoffee) GetDescription() string {
	return "Simple Coffee"
}

type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) SetCoffee(coffee Coffee) {
	d.coffee = coffee
}

func (d *CoffeeDecorator) GetCost() float64 {
	return d.coffee.GetCost()
}

func (d *CoffeeDecorator) GetDescription() string {
	return d.coffee.GetDescription()
}

type MilkCoffee struct {
	CoffeeDecorator
}

func (m *MilkCoffee) GetCost() float64 {
	return m.coffee.GetCost() + 2.0
}

func (m *MilkCoffee) GetDescription() string {
	return m.coffee.GetDescription() + ", Milk"
}

type VanillaCoffee struct {
	CoffeeDecorator
}

func (v *VanillaCoffee) GetCost() float64 {
	return v.coffee.GetCost() + 3.0
}

func (v *VanillaCoffee) GetDescription() string {
	return v.coffee.GetDescription() + ", Vanilla"
}

func main() {
	someCoffee := &SimpleCoffee{}

	fmt.Println(someCoffee.GetCost())         // 10.0
	fmt.Println(someCoffee.GetDescription())  // Simple Coffee

	milkCoffee := &MilkCoffee{}
	milkCoffee.SetCoffee(someCoffee)

	fmt.Println(milkCoffee.GetCost())         // 12.0
	fmt.Println(milkCoffee.GetDescription())  // Simple Coffee, Milk

	vanillaMilkCoffee := &VanillaCoffee{}
	vanillaMilkCoffee.SetCoffee(milkCoffee)

	fmt.Println(vanillaMilkCoffee.GetCost())		// 15.0
	fmt.Println(vanillaMilkCoffee.GetDescription()) // Simple Coffee, Milk, Vanilla
}