package main

import (
	"fmt"
	"reflect"
)

// var (
// 	name   string
// 	module float64
// )
//If outside function then we need to use var
var (
	name, module = "kartik", 2.6
)

func main() {

	//inside the function, no need to use var
	new_name := "rao"
	module_name := "go_lang"

	ptr := &module_name //init a pointer

	fmt.Println("Value of name is ", name, " and type is ", reflect.TypeOf(name))
	fmt.Println("Value of new_name is ", new_name, " and type is ", reflect.TypeOf(new_name))

	fmt.Println("Value of module is", module, " and type is ", reflect.TypeOf(module))
	fmt.Println("Value of module_name is", module_name, " and type is ", reflect.TypeOf(module_name))

	//printing memory address of the variables

	fmt.Println("Address of name is ", &name)
	fmt.Println("Address of module is ", &module)

	fmt.Println("Address of name is ", &new_name)
	fmt.Println("Address of module is ", &module_name)

	fmt.Println("ptr value and address", *ptr, ptr)
}
