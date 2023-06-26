## Program structure ##
> 01-hello-world.go
```
/* package declaration */
package main

/* import dependency packages */

/* package level variable / type declarations */

/* package init function */

/* main function */
func main() {
	print("Hello World!")
}

/* other functions */
```

## To run the program ##
- go run <filename.go>

## To create a build ##
- go build <filename.go>

## To create a build in a different name ##
- go build -o <binary_name> <filename.go>

## To get the list of all the env variables used by the go tools ##
- go env

## To get the specific env variables ##
- go env <var_name_1> <var_name_2>

## Env variables for cross compilation ##
- GOARCH
- GOOS

## TO change the env variable values ##
- go env -w <var_1 = val_1> <var_2 = val_2>

## TO get the list of GOARCH & GOOS values ##
- go tool dist list

## To create a build for different platform ##
-  GOARCH=386 GOOS=windows go build 01-hello-world.go 

# Data Types #
- string
- bool
- Integer family
	- int
	- int8
	- int16
	- int32
	- int64
- Unsigned Integer family
	- uint
	- uint8
	- uint16
	- uint32
	- uint64
- Float family
	- float32
	- float64
- Complex number family
	- complex64 ( real [float32], imaginary [float32])
	- complex128 ( real [float64], imaginary [float64])
- Type aliases
	- byte (alias for uint8)
	- rune (alias for int32) (unicode code point)

## Variable Declarations ##
### function scope ###
- Can use ":=" in a function
- Cannot have unused variables in a function

### package scope ###
- Cannot use ":=" at package level
- Can have unused variables at package level