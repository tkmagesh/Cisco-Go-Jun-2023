package main

import "fmt"

func main() {
	// var nos []int // NO memory allocation & initialized
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos []int = []int{3, 1, 4}

	//type inference
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	fmt.Println("Appending new values")
	nos = append(nos, 10)
	nos = append(nos, 20, 30)
	fmt.Println(nos)

	fmt.Println("Appending another slice values")
	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)
	//
	//

	fmt.Println("Iterating using indexer")
	for i := 0; i < len(nos); i++ {
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

	fmt.Println("slicing")
	fmt.Println("nos[2:5] =", nos[2:5])
	fmt.Println("nos[2:] =", nos[2:])
	fmt.Println("nos[:5] =", nos[:5])

	subset := nos[:5]
	fmt.Printf("len(subset) = %d, cap(subset) = %d, subset = %v\n", len(subset), cap(subset), subset)

	fmt.Println("Creating a copy")
	nosCopy := make([]int, 5, 5)
	copy(nosCopy, nos)
	nosCopy[0] = 9999
	fmt.Printf("nos = %v, nosCopy = %v\n", nos, nosCopy)
}
