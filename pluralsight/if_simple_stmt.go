package main

import (
	"fmt"
)

func main() {

	// a and b is initialized in the if statement only
	// and it has scope in if, else-if and else, after that its garbage collected
	if a, b := 10, 100; a < b {
		fmt.Println("b is greater than a")
	} else if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("Yeradu same")
	}

}
