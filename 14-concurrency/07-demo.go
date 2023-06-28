package main

import (
	"fmt"
)

//share memory by communicating (using channels)

/*
func main() {
	var ch chan int
	ch = make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(ch, wg, 100, 200)
	result := <-ch
	wg.Wait() //
	fmt.Println(result)
}

func add(ch chan int, wg *sync.WaitGroup, x, y int) {
	result := x + y
	ch <- result
	wg.Done()
}
*/

func main() {
	ch := make(chan int)
	go add(ch, 100, 200)
	result := <-ch
	fmt.Println(result)
}

func add(ch chan int, x, y int) {
	result := x + y
	ch <- result
}
