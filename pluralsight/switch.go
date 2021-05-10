package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//case has implicit break statements
	switch "docker" { //here we can give expression also
	case "linux":
		fmt.Println("\n Here are some linux courses")
	case "windows":
		fmt.Println("\n Here are some Windows courses")
	case "docker":
		fmt.Println("\n Here are some docker courses")
		fallthrough // means that docker and kubernetes also gets evaluated
	case "kubernetes":
		fmt.Println("\n Here are some kubernetes courses")
	default:
		fmt.Println("\n This is not part of any course")
	}

	switch tmeVal := random(); tmeVal {
	case 0, 2, 4, 6, 8:
		fmt.Println("\nGot EVEN number", tmeVal)
	case 1, 3, 5, 7, 9:
		fmt.Println("\nGot ODD number", tmeVal)
	}

}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10) //return number between 0 and 9
}
