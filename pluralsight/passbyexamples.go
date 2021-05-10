package main

import "fmt"

func main() {
	course := "go lang"
	name := "kartik"

	fmt.Println("\n Hi ", name, "Your Course is ", course)

	changeCourse(course)

	fmt.Println("\n Updated course is ", course)

	changeCourseByreference(&course)

	fmt.Println("\n After pass by reference", course)
}

func changeCourse(course string) string {
	course = "data structures"
	fmt.Println("\n Course is now : ", course)
	return course
}

func changeCourseByreference(course *string) string {
	*course = "data structures and algo"
	fmt.Println("\n Course is ", *course)
	return *course
}
