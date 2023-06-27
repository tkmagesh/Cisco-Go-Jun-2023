package main

import "fmt"

func main() {
	/*
		var product struct {
			id   int
			name string
			cost float32
		}
		product.id = 100
		product.name = "Pen"
		product.cost = 10
	*/

	/*
		var product struct {
			id   int
			name string
			cost float32
		} = struct {
			id   int
			name string
			cost float32
		}{
			id:   100,
			name: "pen",
			cost: 10,
		}
	*/

	var product = struct {
		id   int
		name string
		cost float32
	}{
		id:   100,
		name: "pen",
		cost: 10,
	}

	// fmt.Printf("%#v\n", product)
	printProduct(product)

}

func printProduct(p struct {
	id   int
	name string
	cost float32
}) {
	fmt.Printf("Product : Id = %d, Name = %q, Cost = %0.2f\n", p.id, p.name, p.cost)
}
