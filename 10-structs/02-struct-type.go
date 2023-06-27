package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float32
}

func main() {
	/*
		var product Product
		product.id = 100
		product.name = "Pen"
		product.cost = 10
	*/

	/*
		var product Product = Product {
			id:   100,
			name: "pen",
			cost: 10,
		}
	*/

	var product = Product{
		id:   100,
		name: "pen",
		cost: 10,
	}

	// fmt.Printf("%#v\n", product)
	printProduct(product)

	var p1 = Product{91, "Apple", 50}
	var p2 = Product{91, "Apple", 50}
	fmt.Printf("%p, %p\n", &p1, &p2)
	fmt.Println(p1 == p2)

	fmt.Println("Before applying discount, product : ", product)
	applyDiscount(&product, 10)
	fmt.Println("After applying discount, product : ", product)

	// var productPtr *Product = &product
	var productPtr = &product
	// fmt.Println((*productPtr).id)
	fmt.Println(productPtr.id)
}

func printProduct(p Product) {
	fmt.Printf("Product : Id = %d, Name = %q, Cost = %0.2f\n", p.id, p.name, p.cost)
}

func applyDiscount(p *Product, discount float32) {
	// (*p).cost = (*p).cost * ((100 - discount) / 100)
	p.cost = p.cost * ((100 - discount) / 100)
}
