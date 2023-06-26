package main

import "fmt"

func main() {
	// const pi float32 = 3.14
	// const pi = 3.14
	// pi = 2

	/*
		const (
			pi          float32 = 3.14
			app_version string  = "1.0"
		)
	*/

	const (
		pi          = 3.14
		app_version = "1.0"
	)
	fmt.Println(pi)
	fmt.Println(app_version)

	//unused constants are allowed
	const x = 100
}
