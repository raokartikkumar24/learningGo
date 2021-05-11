package main

import (
	"fmt"
)

func main() {
	//slice will not create values... kind of pointers
	//underneath slice is array and manipulating the slice updates the array as well
	//myCourses := make([]string, 5, 10) // type, length and capacity
	myCourses := []string{"one", "two", "three"}
	fmt.Printf("Length = %d \n Capacity = %d \n", len(myCourses), cap(myCourses))

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Value at 1 ", mySlice[1])

	mySlice[2] = 100
	fmt.Println("Value at 2 ", mySlice[2])

	//slice of slice
	sliceofslice := mySlice[2:5] //include 2 until 4 , 5 is not included
	fmt.Println(sliceofslice)

	sliceofslice[0] = 1000 // this will change the value of mySlice also since this is a reference

	fmt.Println(mySlice)

	append_slice := make([]int, 5, 10)

	fmt.Printf("\n Initial Length = %d \n Capacity = %d \n", len(append_slice), cap(append_slice))

	for i := 1; i < 17; i++ {
		append_slice = append(append_slice, i)
		//capacity will keep on increasing
		fmt.Printf("\nCapacity = %d", cap(append_slice))
	}

	for _, i := range append_slice {
		fmt.Println(i)
	}

	//appending slice to slice
	one_Slice := []int{1, 2, 3}
	fmt.Println("original one", one_Slice)

	two_slice := []int{100, 200, 300}
	one_Slice = append(one_Slice, two_slice...)
	fmt.Println("udpated one ", one_Slice)
	fmt.Println(two_slice)
}
