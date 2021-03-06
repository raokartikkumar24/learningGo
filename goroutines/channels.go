package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) //creating a channel to communicate between 2 goroutines
	//this channel can receive and send int

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("received", <-ch)
		fmt.Println("sending 24")
		ch <- 24
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("Sending 42")
		ch <- 42 //writting to the channel
		//sleep here for some time otherwise there will deadlock as we are doing sending and receiving in the same routine
		time.Sleep(5 * time.Millisecond)
		fmt.Println("received ", <-ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()

}
