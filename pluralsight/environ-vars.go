package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Hostname())
	fmt.Println(os.Environ())

}
