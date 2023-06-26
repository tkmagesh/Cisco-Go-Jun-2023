package main

import "fmt"

func main() {
	/*
		var name string
		fmt.Println("Enter your name :")
		fmt.Scanln(&name)
		fmt.Println(name)
	*/

	var fname, lname string
	fmt.Println("Enter your firstname & lastname")
	fmt.Scanln(&fname, &lname)
	fmt.Printf("%s %s\n", fname, lname)

	/*
		fmt.Println("Enter your lastname, firstname")
		fmt.Scanf("%s, %s\n", &lname, &fname)
		fmt.Println(fname, lname)
	*/
}
