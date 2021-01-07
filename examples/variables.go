package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

var (
	pkg = "Variables.go" //Package level variable can be used in multiple functions
)

func main() {

	name := strings.Title(os.Getenv("USERNAME")) //Name of subscriber
	course := "Docker Deep Dive"                 //Course being viewed
	module := 3.2
	ptr := &module //References a pointer
	fmt.Println("Welcome to", pkg)
	fmt.Println("\nName is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))

	fmt.Println("Memory address of *module* variable is", ptr)
	fmt.Println("Pointer value of *module* variable is", *ptr) //De-references a pointer

	fmt.Println("\nHi", name, "you are currently watching", course)
	tryChangeCourse(course)
	fmt.Println("\nYou are now watching", course)

	fmt.Println("\nHi", name, "you are currently watching", course)
	changeCourse(&course)
	fmt.Println("\nYou are now watching", course)

}

func tryChangeCourse(course string) string {

	course = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", course)

	return course
}

func changeCourse(course *string) string {

	*course = "First Look: Native Docker Clustering"
	fmt.Println("\nTrying to change your course to", *course)
	return *course
}
