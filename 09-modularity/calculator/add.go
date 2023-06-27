package calculator

import "fmt"

func init() {
	fmt.Println("calculator init invoked [add.go]")
}

func Add(x, y int) int {
	opCount["Add"]++
	return x + y
}
