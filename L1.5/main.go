package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func reader(dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range dataCh {
		fmt.Println(v)
	}
}

func writer(dataCh chan<- int, done <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			close(dataCh)
			return
		default:
			dataCh <- rand.Intn(100)
		}
	}
}

func main() {
	timeout := time.After(2 * time.Second)
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go reader(ch, &wg)
	go writer(ch, timeout, &wg)
	wg.Wait()
}
