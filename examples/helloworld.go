package main

import (
	"fmt"
	"strconv"
)

func num() {
	var num1 = 5
	var rates float32 = 4.5
	fmt.Println("num1 is " +
	strconv.Itoa(num1) +
	" and rates is " + 
	strconv.FormatFloat(
	float64(rates), 'f',2,32))
}

func main() {
	fmt.Println("Hello world")
	// Calling num function
	num()
}
