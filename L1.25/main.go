package main

import (
	"fmt"
	"time"
)

func Sleeppy(duration time.Duration) {
	done := make(chan struct{})
	go func() {
		t := time.NewTimer(duration)
		<-t.C
		close(done)
	}()
	<-done
}

func main() {
	fmt.Println("Start and will wait 2 seconds")
	Sleeppy(2 * time.Second)
	fmt.Println("End")
}
