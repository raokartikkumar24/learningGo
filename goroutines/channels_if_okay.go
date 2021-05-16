package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok { //incase the sending channel is closed then ok will be false and msg will be 0
			//but if the msg itself is 0 then ok becomes true so we can know if its legitimate or an error
			//printing only if the ok is true
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 0
		//close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()

}
