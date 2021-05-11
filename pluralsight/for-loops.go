package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {

		if timer == 0 {
			fmt.Println("Boom!!!!")
			break
		}

		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	coursesNeeded := []string{"go pgm", "malware analysis", "reverse engineering"}

	//for index, i := range coursesNeeded
	for _, i := range coursesNeeded { //here we are ignoring the index so we write '_'
		fmt.Println(i)
	}
}
