package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//without waitgrp main will not wait for the other go routines to finishe its like thread.join()
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	//adding 'go' keyword will spawn a new go routine
	// if we remove the 'go' keyword then its just anoymous function
	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("PLuralsight")
	}()

	waitGrp.Wait()

}
