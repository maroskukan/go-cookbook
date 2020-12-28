package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Flag package string function parses command line arguments
	path := flag.String("path", "myapp.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL")

	flag.Parse()

	// Opening a resource
	f, err := os.Open(*path)
	// Error handling
	if err != nil {
		log.Fatal(err)
	}
	// Resource cleanup
	defer f.Close()
	// Reader for reading file content
	r := bufio.NewReader(f)
	// Looping through the file
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		// Filter functionality
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
