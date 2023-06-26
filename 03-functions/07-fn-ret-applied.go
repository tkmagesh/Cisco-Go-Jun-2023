package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//v1
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	//v3
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	//v4
	// logging
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)

		logMultiply := getLogOperation(func(x, y int) {
			fmt.Println("Multiply Result :", x*y)
		})
		logMultiply(100, 200)
	*/

	//profile
	/*
		profiledAdd := getProfiledOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := getProfiledOperation(subtract)
		profiledSubtract(100, 200)
	*/

	//logging & profiling
	/*
		profiledAdd := getProfiledOperation(add)
		logProfiledAdd := getLogOperation(profiledAdd)
		logProfiledAdd(100, 200)

		profiledSubtract := getProfiledOperation(subtract)
		logProfiledSubtract := getLogOperation(profiledSubtract)
		logProfiledSubtract(100, 200)
	*/

	getLogOperation(getProfiledOperation(add))(100, 200)
	getLogOperation(getProfiledOperation(subtract))(100, 200)
	getLogOperation(getProfiledOperation(func(x, y int) {
		fmt.Println("Multiply result :", x*y)
	}))(100, 200)
}

//v5
func getProfiledOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

//v4
func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("operation started")
		operation(x, y)
		log.Println("operation completed")
	}
}

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
