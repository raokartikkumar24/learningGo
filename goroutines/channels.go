package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) //creating a channel to communicate between 2 goroutines
	//this channel can receive and send int

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("received")
		fmt.Println(<-ch) //receving '<-'
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("Sending 42")
		ch <- 42 //writting to the channel
		wg.Done()
	}(ch, wg)

	wg.Wait()

}
