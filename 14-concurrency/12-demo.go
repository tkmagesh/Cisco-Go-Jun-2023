/* Using channels for streaming data */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := fn()
	/* go func() {
		ch <- 100
	}() */
	for val := range ch {
		fmt.Println(val)
	}
}

// producer
func fn() <-chan int {
	ch := make(chan int)
	go func() {
		count := rand.Intn(20)
		for i := 1; i <= count; i++ {
			ch <- i * 10
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Closing the channel")
		close(ch)
	}()
	return ch
}
