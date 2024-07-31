package main

import (
	"flag"
	"fmt"
)

func main() {
	verbose := flag.Bool("v", false, "enable verbose mode")
	name := flag.String("name", "Vishnu", "Enter your name")

	flag.Parse()

	if *verbose {
		fmt.Println("Verbose mode enabled")
	}

	fmt.Printf("Hello %s", *name)

}
