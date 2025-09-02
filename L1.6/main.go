package main

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

var stop atomic.Bool

func CancelWithRuntime(){
	defer func() {
		fmt.Println("With runtime goroutine canceled")
	}()
	runtime.Goexit()
}
func CancelWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context goroutine canceled")
			return
		default:
			fmt.Println("Doing work with context...")
			time.Sleep(time.Second)
		}
	}
}

func CancelWithAtomic(){
	defer func() {
		fmt.Println("With atomic goroutine canceled")
	}()
	
	for !stop.Load(){
		fmt.Println("Doing smth with atomic")
		time.Sleep(time.Second)
	}
	
}

func CancelWithDoneChannel(done chan struct{}){
	for {
		select {
			case <-done:
				fmt.Println("Done channel goroutine canceled")
				return
			default:
				fmt.Println("Doing work with done channel...")
				time.Sleep(time.Second)
		}
	}
}

func CancelWithTimeAfter(doneWithTime <-chan time.Time){
	for {	
		select {
			case <-doneWithTime:
				fmt.Println("TimeAfter goroutine canceled")
				return
			default:
				fmt.Println("Doing work with time after...")
				time.Sleep(time.Second)
		}
	}
}

func CancelWithNotification(doneWithNotification <- chan time.Time, other <- chan string){
	select {
	case <-doneWithNotification:
		fmt.Println("Goroutine canceled with notification")
		return
	case <-other:
		fmt.Println("Received message from other")
	default:
		fmt.Println("Doing work with notification...")
	}
}

func CloseChan(channel <- chan string){
	defer func(){
		fmt.Println("Goroutine finished because channel closed")
	}()
	for data := range channel {
		fmt.Println("Received message from channel:", data)
	}
}

func JustFinished(){
	fmt.Println("Just finished")
}

func CancelWithPanic(){
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Panic occurred")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	
	chanToClose := make(chan string)
	chanWithMessage := make(chan string)
	
	
	timeout := time.After(2 * time.Second)
	timeoutForSecond := time.After(2 * time.Second)
	go CloseChan(chanToClose)
	chanToClose <- "Hello"
	go CancelWithTimeAfter(timeout)
	go CancelWithNotification(timeoutForSecond, chanWithMessage)
	chanWithMessage <- "Hello"
	go JustFinished()
	go CancelWithContext(ctx)
	go CancelWithRuntime()
	go CancelWithAtomic()
	go CancelWithPanic()
	channelForDone := make(chan struct{})
	go CancelWithDoneChannel(channelForDone)
	channelForDone <- struct{}{}
	time.Sleep(time.Second * 3)
	
	stop.Store(true)
	cancel()
	close(chanToClose)
	
	time.Sleep(4*time.Second)
}