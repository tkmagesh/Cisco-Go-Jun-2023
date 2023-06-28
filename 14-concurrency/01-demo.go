package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduled to be executed in future
	f2()
	// current function execution is blocked
	time.Sleep(time.Second)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
	/*
		for {

		}
	*/
}
