/*
Write a function "getPrimes" that generates prime numbers within the given input (start & end)
and print them
*/

package main

import "fmt"

func main() {
	primes := getPrimes(3, 100)
	fmt.Println(primes)
}

func getPrimes(start, end int) []int {
	result := make([]int, 0)
	for i := start; i <= end; i++ {
		if isPrime(i) {
			result = append(result, i)
		}
	}
	return result
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
