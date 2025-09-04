package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Incrementer struct {
	Counter int64
}

func increment(x *Incrementer, wg *sync.WaitGroup) {
	wg.Done()
	for i := 0; i < 10; i++ {
		atomic.AddInt64(&x.Counter, 1)
	}
}

func main() {
	var x Incrementer
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go increment(&x, &wg)
	}
	wg.Wait()
	fmt.Println(x.Counter)
}
