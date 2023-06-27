package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	[deferred] - main")
	}()
	fmt.Println("main started")
	result := f1()
	fmt.Println("f1 result :", result)
	fmt.Println("main completed")
}

func f1() (result int) {

	defer fmt.Println("	[deferred] - f1(1)")
	defer fmt.Println("	[deferred] - f1(2)")
	defer func() {
		fmt.Println("	[deferred] - f1(3)")
		result = 200
	}()
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
	result = 100
	return
}

func f2() {
	defer func() {
		fmt.Println("	[deferred] - f2")
	}()
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
