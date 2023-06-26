package main

import (
	"fmt"
	"log"
)

func main() {
	//v1
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	//v2
	/*
		log.Println("operation started")
		add(100, 200)
		log.Println("operation completed")

		log.Println("operation started")
		subtract(100, 200)
		log.Println("operation completed")
	*/

	//v3
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	//v4
	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(x, y int) {
		fmt.Println("Multiply Result :", x*y)
	}, 100, 200)
}

//v4
func logOperation(operation func(int, int), x, y int) {
	log.Println("operation started")
	operation(x, y)
	log.Println("operation completed")
}

// v3
func logAdd(x, y int) {
	log.Println("operation started")
	add(x, y)
	log.Println("operation completed")
}

func logSubtract(x, y int) {
	log.Println("operation started")
	subtract(100, 200)
	log.Println("operation completed")
}

//v1
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
