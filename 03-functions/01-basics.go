package main

import "fmt"

func main() {
	/* no parameters, no return values */
	sayHi()

	/* 1 parameter, no return values */
	greet("Magesh")

	/* 1 parameter, 1 return value */
	print(getGreetMsg("Suresh"))

	/* 2 parameters, 1 return value */
	fmt.Println("add(100,200) =", add(100, 200))

	// 2 parameters, 2 return values
	// fmt.Println(divide(100, 7))

	/*
		q, r := divide(100, 7)
		fmt.Printf("dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/

	//Printing ONLY quotient
	// q, r := divide(100, 7) // error : r is declared but not used
	// q := divide(100, 7) // error : assignment mismatch: 1 variable but divide returns 2 values

	q, _ := divide(100, 7)
	fmt.Printf("dividing 100 by 7, quotient = %d \n", q)

	//Printing ONLY remainder
	_, r := divide(100, 7)
	fmt.Printf("dividing 100 by 7, remainder = %d \n", r)

	fmt.Println("Using named results")
	fmt.Println(divide(10, 10))

}

func sayHi() {
	fmt.Println("Hi")
}

func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
}

/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// named results

/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
