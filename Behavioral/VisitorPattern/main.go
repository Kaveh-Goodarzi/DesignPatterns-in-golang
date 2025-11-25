package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Accept(visitor ShapeVisitor)
}

type ShapeVisitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
	VisitSquare(square *Square)
}

type Circle struct {
	radius float64
}

func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

func (c *Circle) getRadius() float64 {
	return c.radius
}

type Square struct {
	length float64
}

func (s *Square) Accept(visitor ShapeVisitor) {
	visitor.VisitSquare(s)
}

func (s *Square) getLength() float64 {
	return s.length
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r)
}

func (r *Rectangle) getWidth() float64 {
	return r.width
}

func (r *Rectangle) getHeight() float64 {
	return r.height
}

type shapeVisitor struct {
	area      float64
	perimeter float64
}

func (v *shapeVisitor) VisitCircle(circle *Circle) {
	v.area = math.Pi * circle.getRadius() * circle.getRadius()
	v.perimeter = 2 * math.Pi * circle.getRadius()
}

func (v *shapeVisitor) VisitRectangle(rectangle *Rectangle) {
	v.area = rectangle.getWidth() * rectangle.getHeight()
	v.perimeter = 2 * (rectangle.getWidth() + rectangle.getHeight())
}

func (v *shapeVisitor) VisitSquare(square *Square) {
	v.area = square.getLength() * square.getLength()
	v.perimeter = 4 * square.getLength()
}

func (v *shapeVisitor) getArea() float64 {
	return v.area
}

func (v *shapeVisitor) getPerimeter() float64 {
	return v.perimeter
}

func main() {
	shapes := []Shape{
		&Circle{radius: 4},
		&Square{length: 5},
		&Rectangle{width: 3, height: 6},
	}

	for i, shape := range shapes {
		visitor := &shapeVisitor{}
		shape.Accept(visitor)

		fmt.Printf("Shape #%d:\n", i+1)
		fmt.Printf("  Area:      %.2f\n", visitor.getArea())
		fmt.Printf("  Perimeter: %.2f\n\n", visitor.getPerimeter())
	}
}
