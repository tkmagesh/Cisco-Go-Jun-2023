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
	/*
		err2 := errors.Unwrap(err1)
		err3 := errors.Unwrap(err2)
		fmt.Println("err 1 :", err1)
		fmt.Println("err 2 :", err2)
		fmt.Println("err 3 :", err3)
	*/
	fmt.Println("err1 :")
	fmt.Println(err1)
	fmt.Println("=================")

	if errors.Is(err1, ErrF1Error) {
		fmt.Println("f1 error occurred")
	}
	if errors.Is(err1, ErrF2Error) {
		fmt.Println("f2 error occurred")
	}
	if errors.Is(err1, ErrF3Error) {
		fmt.Println("f3 error occurred")
	}
}

func f1() error {
	err := f2()
	return errors.Join(ErrF1Error, err)
}

func f2() error {
	err := f3()
	return errors.Join(ErrF2Error, err)
}

func f3() error {
	return ErrF3Error
}
