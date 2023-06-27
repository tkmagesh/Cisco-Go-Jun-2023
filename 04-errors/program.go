package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

var ErrDivideByZeroErro error = errors.New("cannot divide by zero")

func main() {
	divisor := 0
	q, r, err := divide(100, divisor)
	if err == ErrDivideByZeroErro {
		fmt.Println("Do not attempt to divide by zero")
		return
	}
	if err != nil {
		fmt.Println("Something went wrong :", err)

		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := DivideByZeroErro
		return 0, 0, err
	}
	quotient, remainder := x/y, x%y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = ErrDivideByZeroErro
		debug.PrintStack()
		return
	}
	quotient, remainder = x/y, x%y
	return
}
