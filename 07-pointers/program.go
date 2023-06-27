package main

import "fmt"

func main() {
	var no int
	no = 100

	var noPtr *int
	noPtr = &no // get the address of "no" variable
	fmt.Println(noPtr, no)

	//dereferencing - accessing the value at an address(pointer)
	valAtNoPtr := *noPtr
	fmt.Println(valAtNoPtr)

	fmt.Println(no == *(&no))
	fmt.Println("Before incrementing : no =", no)
	increment(&no)
	fmt.Println("After incrementing : no =", no)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)

}

func increment(x *int) {
	fmt.Println(x)
	*x++
}

func swap(x, y *int) /* do not return anything */ {
	*x, *y = *y, *x
}
