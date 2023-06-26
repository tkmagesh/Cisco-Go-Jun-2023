package main

import "fmt"

func main() {
	var x int32 = 100
	var f float64 = float64(x) // use the typname like a function for type conversion
	fmt.Println(f)
}
