package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Occaecat proident amet non eu cupidatat. Sunt reprehenderit ipsum ullamco ullamco tempor proident excepteur eiusmod eu. Eiusmod minim aliquip tempor ad minim irure dolor nisi nulla quis aute Lorem. Esse culpa id magna magna nulla laboris aute minim consectetur et incididunt."
	x = true
	x = 99.99
	fmt.Println(x)

	x = 100
	// x = true
	// y := x.(int) + 200
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// x = 100
	// x = "Occaecat proident amet non eu cupidatat. Sunt reprehenderit ipsum ullamco ullamco tempor proident excepteur eiusmod eu. Eiusmod minim aliquip tempor ad minim irure dolor nisi nulla quis aute Lorem. Esse culpa id magna magna nulla laboris aute minim consectetur et incididunt."
	// x = true
	// x = 99.99
	/*
		x = struct {
			Id   int
			Name string
			Cost float32
		}{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/
	x = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, 99% of x =", val*0.99)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	/*
		case struct {
			Id   int
			Name string
			Cost float32
		}:
			fmt.Println("x is a product", val.Id, val.Name, val.Cost)
	*/
	case Product:
		fmt.Println("x is a product", val.Id, val.Name, val.Cost)
	default:
		fmt.Println("x is an unknown type")
	}

}
