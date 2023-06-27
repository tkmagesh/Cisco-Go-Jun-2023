package main

import (
	"errors"
	"fmt"
)

var ErrF1Error = errors.New("error from f1")
var ErrF2Error = errors.New("error from f2")
var ErrF3Error = errors.New("error from f3")

func main() {
	err1 := f1()
	err2 := errors.Unwrap(err1)
	err3 := errors.Unwrap(err2)
	fmt.Println("err 1 :", err1)
	fmt.Println("err 2 :", err2)
	fmt.Println("err 3 :", err3)
}

func f1() error {
	err := f2()
	return fmt.Errorf("%w : %w", ErrF1Error, err)
}

func f2() error {
	err := f3()
	return fmt.Errorf("%w : %w", ErrF2Error, err)
}

func f3() error {
	return ErrF3Error
}
