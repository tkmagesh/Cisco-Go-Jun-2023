package main

import "fmt"

func main() {
	var c1 complex128 = 4 + 5i
	fmt.Println(c1)
	c2 := 6 + 7i
	fmt.Println(c2)
	c3 := c1 + c2
	fmt.Printf("%v + %v = %v\n", c1, c2, c3)
	fmt.Printf("real(c3) = %v, imag(c3) = %v\n", real(c3), imag(c3))
}
