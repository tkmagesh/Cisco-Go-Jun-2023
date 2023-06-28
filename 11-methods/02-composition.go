package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float32
}

func (p Product) FormatProduct() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.id, p.name, p.cost)
}

// Behavior as a method (function with receiver)
func (p *Product) ApplyDiscount(discount float32) {
	p.cost = p.cost * ((100 - discount) / 100)
}

type Dummy struct {
	name string
}

type PerishableProduct struct {
	Product
	// Dummy
	cost   float32 // overriding the Product.cost
	expiry string
}

// overriding the Product.FormatProduct() method
func (p PerishableProduct) FormatProduct() string {
	// return fmt.Sprintf("%s, Expiry = %q", p.Product.FormatProduct(), p.expiry)
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Expiry = %q", p.id, p.name, p.cost, p.expiry)
}

func main() {
	/*
		grapes := PerishableProduct{
			Product{100, "Grapes", 50},
			"2 Days",
		}
	*/

	grapes := PerishableProduct{
		Product: Product{
			id:   100,
			name: "Grapes",
			cost: 50,
		},
		expiry: "2 Days",
		cost:   90,
	}

	// accessing the methods of the base (Product) through the derived type instance (Perishable Product)
	fmt.Println(grapes.FormatProduct())
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.FormatProduct())
}
