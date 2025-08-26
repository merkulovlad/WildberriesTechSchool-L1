package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("%d worker %d\n", val, i)
	}
}
func main() {
	var n int
	var wg sync.WaitGroup
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
		go worker(i, ch, &wg)
	}
	wg.Wait()

}
