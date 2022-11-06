package main

import (
	"fmt"
	"os"
)

func main() {

	c := len(os.Args) - 1
	if c < 1 {
		fmt.Fprintf(os.Stderr, "[usadge] %s <name>", os.Args[0])
		os.Exit(1)
	}

	greet := MakeHello(os.Args[1])
	fmt.Printf(greet + "\n")
}

func MakeHello(name string) string {
	return "Hello World! " + name
}
