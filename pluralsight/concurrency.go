package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0

func main() {

	//without waitgrp main will not wait for the other go routines to finishe its like thread.join()
	var waitGrp sync.WaitGroup
	var _mutex sync.Mutex

	waitGrp.Add(4)

	go func() {
		thread1(&_mutex)
		waitGrp.Done()
	}()

	go func() {
		thread2(&_mutex)
		waitGrp.Done()
	}()

	//adding 'go' keyword will spawn a new go routine
	// if we remove the 'go' keyword then its just anonymous function
	go func() {
		//defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
		waitGrp.Done()
	}()

	go func() {
		//defer waitGrp.Done()
		fmt.Println("PLuralsight")
		waitGrp.Done()
	}()

	waitGrp.Wait()

}

func thread1(m *sync.Mutex) {
	for i := 0; i < 100; i++ {
		m.Lock()
		count = count + 1
		fmt.Println("Thread 1, count = ", count)
		m.Unlock()

	}
}

func thread2(m *sync.Mutex) {
	for i := 0; i < 100; i++ {
		m.Lock()
		count = count - 1
		fmt.Println("Thread 2, count = ", count)
		m.Unlock()

	}
}
