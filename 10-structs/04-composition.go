package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float32
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
	fmt.Println(grapes)
	//accessing the attributes
	fmt.Println(grapes.expiry)
	fmt.Println(grapes.Product.id, grapes.Product.name)
	fmt.Println(grapes.id, grapes.name)
	fmt.Println(grapes.cost)
}
