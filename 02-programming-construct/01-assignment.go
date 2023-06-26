/* Print all the prime numbers between 3 to 100 */
package main

import "fmt"

func main() {
OUTER_LOOP:
	for no := 3; no <= 100; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue OUTER_LOOP
			}
		}
		fmt.Printf("%d is a prime no\n", no)
	}
}
