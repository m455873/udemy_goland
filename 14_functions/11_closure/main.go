package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}


func main() {
	increment := wrapper()
	fmt.Println(increment()) // returns 1
	fmt.Println(increment()) // returns 2

	// the variable scope of x becomes part of main, namely the increment instance... as long as this is in scope it will remember the value of x rather than restarting from 0
}
