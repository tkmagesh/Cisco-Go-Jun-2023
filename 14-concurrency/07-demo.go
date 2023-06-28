package main

import (
	"fmt"
	"sync"
)

//share memory by communicating (using channels)

func main() {
	var ch chan int
	ch = make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(ch, wg, 100, 200)
	wg.Wait()
	result := <-ch
	fmt.Println(result)
}

func add(ch chan int, wg *sync.WaitGroup, x, y int) {
	defer wg.Done()
	result := x + y
	ch <- result
}
