/* functions as values to variables */

package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var getGreetMsg func(string) string
	getGreetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!", userName)
	}
	fmt.Println(getGreetMsg("Suresh"))

	var operation func(int, int) int
	operation = func(x, y int) int {
		return x + y
	}
	fmt.Println(operation(100, 200))

	operation = func(x, y int) int {
		return x - y
	}
	fmt.Println(operation(100, 200))

	operation = func(x, y int) int {
		return x * y
	}
	fmt.Println(operation(100, 200))

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
