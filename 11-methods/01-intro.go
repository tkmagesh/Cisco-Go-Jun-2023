package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float32
}

func main() {
	/*
		pen := Product{
			id:   100,
			name: "Pen",
			cost: 10,
		}
	*/

	pen := &Product{
		id:   100,
		name: "Pen",
		cost: 10,
	}

	// fmt.Println(FormatProduct(pen))
	fmt.Println(pen.FormatProduct())

	fmt.Println("After applying 10% discount")
	// ApplyDiscount(&pen, 10)
	// (&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10)
	// fmt.Println(FormatProduct(pen))
	fmt.Println(pen.FormatProduct())
}

/*
func FormatProduct(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.id, p.name, p.cost)
}
*/

func (p Product) FormatProduct() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.id, p.name, p.cost)
}

/*
func ApplyDiscount(p *Product, discount float32) {
	p.cost = p.cost * ((100 - discount) / 100)
}
*/

func (p *Product) ApplyDiscount(discount float32) {
	p.cost = p.cost * ((100 - discount) / 100)
}
