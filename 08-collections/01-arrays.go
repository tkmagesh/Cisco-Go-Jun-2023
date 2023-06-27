package main

import "fmt"

func main() {
	// var nos [5]int // memory allocation & initialized with the default (zero) value of int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}

	//type inference
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	for i := 0; i < 5; i++ {
		nos[i] = i * 10
	}
	fmt.Println("Iterating using indexer")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	newNos := nos
	newNos[0] = 10000
	fmt.Println("newNos =", newNos)
	fmt.Println("nos =", nos)
}
