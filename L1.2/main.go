package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	i = i * i
	fmt.Printf("%d\n", i)
}

func main() {
	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10}
	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go worker(numbers[i], &wg)
	}
	wg.Wait()
}
