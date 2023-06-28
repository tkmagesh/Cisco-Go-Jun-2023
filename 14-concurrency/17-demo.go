/* write a goroutine that generates n fibonocci series where n is given as an input */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generateFib()
	for val := range ch {
		fmt.Println(val)
	}
}
func generateFib() chan int {
	ch := make(chan int)
	stopCh := time.After(10 * time.Second)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

//replaced with time.After()
/*
func timeout(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		time.Sleep(d)
		ch <- time.Now()
	}()
	return ch
}
*/
