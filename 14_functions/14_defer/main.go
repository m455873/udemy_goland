package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	// Prints hello world
	// Useful for closing file opened at point you opened it so it is closed before main exits
	defer world() // delays execution of function to just before main exits
	hello()
}
