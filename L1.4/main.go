package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(i int, ch <-chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: canceled\n", i)
			return
		case val, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("%d worker %d\n", val, i)
		}
	}
}
func main() {
	var n int
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	fmt.Scan(&n)
	ch := make(chan int)
	go func(chan int) {
		for i := 0; i < 1000; i++ {
			ch <- i
			time.Sleep(5 * time.Millisecond)
		}
		close(ch)
	}(ch)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(i, ch, &wg, ctx)
	}
	wg.Wait()

}
