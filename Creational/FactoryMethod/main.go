package main

import "fmt"

// iGun is an interface for gun -> iGun.go
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Gun is a class for gun -> gun.go
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}


// AK-47 is a concrete gun -> ak47.go
type AK47 struct {
	Gun
}

func newAk47() IGun {
	return &AK47{
		Gun: Gun{
			name: "AK-47 Gun",
			power: 4,
		},
	}
}


// Musket is a concrete gun -> musket.go
type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name: "Musket Gun",
			power: 1,
		},
	}
}


// -> gunFactory.go
func getGun(gunType string) (IGun, error) {
	if gunType == "AK-47" {
		return newAk47(), nil
	}

	if gunType == "Musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed!")
}


// Client -> main.go
func main() {
	ak47, _ := getGun("AK-47")
	musket, _ := getGun("Musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
}