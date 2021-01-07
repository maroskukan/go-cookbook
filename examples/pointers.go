package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Maros" //Name of subscriber
	//course := "Docker Deep Dive" //Course being viewed
	module := 3.2
	ptr := &module
	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))

	fmt.Println("Memory address of *module* variable is", ptr)
	fmt.Println("Pointer value of *module* variable is", *ptr)

}
