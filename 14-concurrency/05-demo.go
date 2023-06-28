package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

func main() {
	wg := &sync.WaitGroup{}
	c := &Counter{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go c.increment(wg)
	}
	wg.Wait()
	fmt.Println(c.count)
}
