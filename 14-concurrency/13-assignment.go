/* write a goroutine that generates n fibonocci series where n is given as an input */

package main

import (
	"fmt"
)

func main() {
	ch := generateFib(20)
	for val := range ch {
		fmt.Println(val)
	}
}
func generateFib(n int) chan int {
	ch := make(chan int)

	go func() {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			ch <- x
			x, y = y, x+y
		}
		close(ch)
	}()
	return ch
}
