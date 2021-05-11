package main

import (
	"fmt"
)

func main() {
	//define a type courseMeta
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	// var DockerDeepDive courseMeta
	// DockerDeepDive := new(courseMeta)

	DockerDeepDive := courseMeta{
		Author: "abc xyz",
		Level:  "advanced",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)

	fmt.Println("\n Author is ", DockerDeepDive.Author)

	DockerDeepDive.Rating = 1

	fmt.Println("\n Rating is ", DockerDeepDive.Rating)

}
