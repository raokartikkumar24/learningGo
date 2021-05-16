package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) //creating a buffered channel with capacity 1
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
		ch <- 24
		wg.Done()
	}(ch, wg)

	wg.Wait()

}
