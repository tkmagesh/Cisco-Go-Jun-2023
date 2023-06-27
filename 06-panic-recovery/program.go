package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZeroError = errors.New("Divide by zero error")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("application panicked : ", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	divisor := 0
	q, r, err := divideClient(100, divisor)
	if err == nil {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
		return
	}
	fmt.Println("error occurred : ", err)
	fmt.Println("Calling another divide function.....")
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = er.(error)
		}
		return
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party library
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(ErrDivideByZeroError)
	}
	fmt.Println("Calculating quotient")
	quotient = x / y
	fmt.Println("Calculating remainder")
	remainder = x % y
	return
}
