package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func generateRandomNumber(channel chan <- int) {
	for i := 0; i < 10; i++ {
		channel <- rand.Intn(100)
	}
	close(channel)
}

func multiplyNumber(channel <-chan int, wg *sync.WaitGroup) {
	for num := range channel {
		fmt.Println(num * 2)
	}
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	channel := make(chan int)

	wg.Add(1)
	go generateRandomNumber(channel)
	go multiplyNumber(channel, &wg)

	wg.Wait()
}