package main

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
}

type ProductCatalog struct {
	products []Product
}

func (pc *ProductCatalog) AddProduct(p Product) {
	pc.products = append(pc.products, p)
}

func (pc *ProductCatalog) GetName() string {
	return "Product Catalog:"
}

func (pc *ProductCatalog) GetPrice() float64 {
	total := 0.0
	for _, p := range pc.products {
		total += p.GetPrice()
	}
	return total
}

type Book struct {
	name  string
	price float64
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) GetPrice() float64 {
	return b.price
}

type Gadget struct {
	name  string
	price float64
}

func (g *Gadget) GetName() string {
	return g.name
}

func (g *Gadget) GetPrice() float64 {
	return g.price
}

type Movie struct {
	name  string
	price float64
}

func (m *Movie) GetName() string {
	return m.name
}

func (m *Movie) GetPrice() float64 {
	return m.price
}

type Song struct {
	name  string
	price float64
}

func (s *Song) GetName() string {
	return s.name
}

func (s *Song) GetPrice() float64 {
	return s.price
}

func main() {
	book := &Book{name: "Go Programming", price: 19.99}
	gadget := &Gadget{name: "Smartphone", price: 699.99}
	movie := &Movie{name: "Inception", price: 14.99}
	song := &Song{name: "ma edame darim", price: 1.29}

	catalog := &ProductCatalog{}
	catalog.AddProduct(book)
	catalog.AddProduct(gadget)
	catalog.AddProduct(movie)
	catalog.AddProduct(song)

	fmt.Println(catalog.GetName())
	fmt.Println("Total Price:", catalog.GetPrice())
}
