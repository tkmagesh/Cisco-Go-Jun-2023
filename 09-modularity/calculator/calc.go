package calculator

import "fmt"

var opCount map[string]int

func init() {
	fmt.Println("calculator init invoked [calc.go] - 1")
	opCount = make(map[string]int)
}

func OpCount() map[string]int {
	return opCount
}
