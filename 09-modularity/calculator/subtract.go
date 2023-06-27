package calculator

import "fmt"

func init() {
	fmt.Println("calculator init invoked [subtract.go]")
}

func Subtract(x, y int) int {
	opCount["Subtract"]++
	return x - y
}
