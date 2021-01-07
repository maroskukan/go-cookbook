package main

import (
	"fmt"
	"os"
)

func main() {

	//Printing all evnrionment variables
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	//Printing specific value based on key, useful when passing arguments to instance
	//docker run .. -e "USER=your-username" -e "PASS=your-password"
	fmt.Println(os.Getenv("USER"))

}
