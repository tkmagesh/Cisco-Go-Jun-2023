/* using waitgroups */

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(50)
	var gCount = flag.Int("count", 0, "Number of goroutines to start")
	flag.Parse()
	wg := &sync.WaitGroup{}
	fmt.Printf("Starting %d goroutines... Hit ENTER to start\n", *gCount)
	fmt.Scanln()
	for i := 1; i <= *gCount; i++ {
		wg.Add(1)
		go fn(i, wg)
	}
	wg.Wait()
	fmt.Println("Hit ENTER to shutdown...")
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
