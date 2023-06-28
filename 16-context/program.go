package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()

	valCtx := context.WithValue(rootCtx, "root-key", "root-val")
	// cancelCtx, cancel := context.WithCancel(rootCtx)
	cancelCtx, cancel := context.WithCancel(valCtx)

	fmt.Println("Hit ENTER to shutdown....")
	wg.Add(1)
	go func() {
		fmt.Scanln()
		cancel()
	}()
	go fn(cancelCtx, wg)
	wg.Wait()
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("[fn] root-key =", ctx.Value("root-key"))
	defer wg.Done()

	fnValCtx := context.WithValue(ctx, "fn-key", "fn-value")
	timeoutCtx, cancel := context.WithTimeout(fnValCtx, 2*time.Second)
	defer func() {
		fmt.Println("exiting fn")
		cancel()
	}()
	wg.Add(1)
	go f1(timeoutCtx, wg)

	wg.Add(1)
	go f2(timeoutCtx, wg)

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[fn] cancel signal received")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("fn invoked")
		}

	}
}

func f1(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("[f1] root-key : ", ctx.Value("root-key"))
	fmt.Println("[f1] fn-key : ", ctx.Value("fn-key"))
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[f1] cancel signal received")
			break LOOP
		default:
			time.Sleep(400 * time.Millisecond)
			fmt.Println("f1 invoked")
		}

	}
}

func f2(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("[f2] root-key : ", ctx.Value("root-key"))
	fmt.Println("[f2] fn-key : ", ctx.Value("fn-key"))
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[f2] cancel signal received")
			break LOOP
		default:
			time.Sleep(200 * time.Millisecond)
			fmt.Println("f2 invoked")
		}

	}
}
