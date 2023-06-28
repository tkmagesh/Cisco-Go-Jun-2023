/* Using channels for streaming data */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go fn(ch)
	for {
		if val, isOpen := <-ch; isOpen {
			fmt.Println(val)
			continue
		}
		break
	}
}

// producer
func fn(ch chan int) {
	count := rand.Intn(20)
	for i := 1; i <= count; i++ {
		ch <- i * 10
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("Closing the channel")
	close(ch)
}
