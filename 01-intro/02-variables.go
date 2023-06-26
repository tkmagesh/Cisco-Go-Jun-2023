package main

import "fmt"

// Cannot use := at package level
// age := 30

// Can have unused variables at package level
var age = 30

func main() {

	/*
		var name string
		name = "Magesh"
		// print(name + "\n")
		// fmt.Println(name)
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	/*
		var name string = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	//type inference
	/*
		var name = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	//using :=
	name := "Magesh" //shortcut for line 28
	name += " Kuppan"
	fmt.Printf("Hi %s, Have a nice day!\n", name)

	//cannot have unused variables
	// age := 30

	// Using multiple variables
	/*
		var x int
		var y int
		var result int
		var str string
		x = 100
		y = 200
		result = x + y
		str = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		result = x + y
		str = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		result = x + y
		str = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	// Declaration & initialization
	/*
		var x int = 100
		var y int = 200
		var result int = x + y
		var str string = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x      int    = 100
			y      int    = 200
			result int    = x + y
			str    string = "Sum of %d and %d is %d\n"
		)
	*/
	/*
		var (
			x, y   int    = 100, 200
			result int    = x + y
			str    string = "Sum of %d and %d is %d\n"
		)
	*/

	/*
		var (
			x, y   = 100, 200
			result = x + y
			str    = "Sum of %d and %d is %d\n"
		)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
	*/

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y

	fmt.Printf(str, x, y, result)
}
