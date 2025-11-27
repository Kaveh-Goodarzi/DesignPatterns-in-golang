package main

import "fmt"

// Train.go: Component

type Train interface {
	Arrive()
	Depart()
	PermitArrival()
}

// Mediator.go

type Mediator interface {
	CanArrive(Train) bool
	NotifyAboutDeparture()
}

// passengerTrain.go

type PassengerTrain struct {
	mediator Mediator
}

func (g *PassengerTrain) Arrive() {
	if !g.mediator.CanArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.NotifyAboutDeparture()
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.Arrive()
}

// freightTrain.go

type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) Arrive() {
	if !g.mediator.CanArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.NotifyAboutDeparture()
}

func (g *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted, arriving")
	g.Arrive()
}

// stationManager.go

type StationManager struct {
	isPlatformFree bool
	trains         []Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (sm *StationManager) CanArrive(train Train) bool {
	if sm.isPlatformFree {
		sm.isPlatformFree = false
		return true
	}
	sm.trains = append(sm.trains, train)
	return false
}

func (sm *StationManager) NotifyAboutDeparture() {
	if sm.isPlatformFree {
		sm.isPlatformFree = true
	}
	if len(sm.trains) > 0 {
		firstTrain := sm.trains[0]
		sm.trains = sm.trains[1:]
		firstTrain.PermitArrival()
	}
}

// client (main.go)

func main() {
	StationManager := NewStationManager()
	PassengerTrain := &PassengerTrain{
		mediator: StationManager,
	}
	FreightTrain := &FreightTrain{
		mediator: StationManager,
	}

	PassengerTrain.Arrive()
	FreightTrain.Arrive()
	PassengerTrain.Depart()
}
