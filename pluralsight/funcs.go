package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "go basics"
	name := "kartik"

	fmt.Println(converter(module, name))

	fmt.Println(bestFinish(9, 8, 1, 2, 4))

	fmt.Println(bestFinish(9, 8, 1, 2, 4, 12, 11, 100))

	fmt.Println(bestFinish(9, 8, 10, 2, 4, 12, 11, 111, 88, 8))
}

func converter(module string, name string) (string, string) {
	module = strings.Title(module)
	name = strings.ToUpper(name)
	return module, name
}

//vairable number of arguments
func bestFinish(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}
