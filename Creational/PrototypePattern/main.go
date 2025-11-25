package main

import "fmt"

// 1) Prototype
// 2) Prototype registry
// 3) Client

type ShapeType int

const (
	CircleType ShapeType = 1
	SquareType ShapeType = 2
)

type Shape interface {
	GetId() ShapeType
	PrintTypeProp()
	Clone() Shape
}

type Circle struct {
	Id            ShapeType
	Radius        int
	Diameter      int
	Cricumference int
}

func NewCircle(radius, diameter, cricumference int) Circle {
	return Circle{CircleType, radius, diameter, cricumference}
}

func (c Circle) GetId() ShapeType {
	return c.Id
}

func (c Circle) Clone() Shape {
	return NewCircle(c.Radius, c.Diameter, c.Cricumference)
}

func (c Circle) PrintTypeProp() {
	fmt.Printf("Circle Properties:\n Radius %d, Diameter %d, Cricumference %d\n", c.Radius, c.Diameter, c.Cricumference)
}

type Square struct {
	Id      ShapeType
	Length  int
	Breadth int
}

func NewSquare (Length, Breadth int) Square {
	return Square{SquareType, Length, Breadth}
}

func (s Square) GetId() ShapeType {
	return s.Id
}

func (s Square) Clone() Shape {
	return NewSquare(s.Length, s.Breadth)
}

func (s Square) PrintTypeProp() {
	fmt.Printf("Square Properties:\n Length %d\n Breadth %d\n", s.Length, s.Breadth)
}

func main() {
	Circle := NewCircle(50, 40, 15)
	Circle.GetId()
	Circle.PrintTypeProp()

	Circle2 := Circle.Clone()
	Circle2.GetId()
	Circle2.PrintTypeProp()

	fmt.Println()
	fmt.Println("========")
	fmt.Println()

	Square := NewSquare(10, 15)
	Square.GetId()
	Square.PrintTypeProp()

	Square2 := Square.Clone()
	Square2.GetId()
	Square2.PrintTypeProp()
}