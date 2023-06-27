package main

import "fmt"

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"pen": 4, "pencil": 1}
	/*
		var productRanks map[string]int = map[string]int{
			"pen":    4,
			"pencil": 1,
		}
	*/
	//type inference
	/*
		var productRanks = map[string]int{
			"pen":    4,
			"pencil": 1,
		}
	*/
	/*
		productRanks := map[string]int{
			"pen":    4,
			"pencil": 1,
		}
	*/

	// var productRanks map[string]int = make(map[string]int)
	var productRanks map[string]int = map[string]int{}
	productRanks["pen"] = 4
	productRanks["pencil"] = 1
	productRanks["marker"] = 3
	productRanks["notepad"] = 2
	fmt.Println(productRanks)

	fmt.Println("# of product ranked =", len(productRanks))

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("Check for the existence of a key")
	nonExistingKey := "scribble-pad"
	if val, exists := productRanks[nonExistingKey]; !exists {
		fmt.Printf("Key - %q doesnot exist\n", nonExistingKey)
	} else {
		fmt.Printf("Rank of %q is %d\n", nonExistingKey, val)
	}

	keyToRemove := "pen"
	delete(productRanks, keyToRemove)
	fmt.Println("After removing pen, productRanks = ", productRanks)

}
