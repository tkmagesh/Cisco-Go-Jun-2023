/* write a goroutine that generates n fibonocci series where n is given as an input */

package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop....")
	go notifyStop(stopCh)
	ch := generateFib(stopCh)
	for val := range ch {
		fmt.Println(val)
	}
}

func notifyStop(ch chan<- struct{}) {
	fmt.Scanln()
	ch <- struct{}{}
	// close(ch)
}

func generateFib(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
				break LOOP
			default:
				ch <- x
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}
