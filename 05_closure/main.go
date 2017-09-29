package main

import "fmt"

var x = 0 // outer package scope

func increment() int {
	x++
	return x
}

func wrapper() func() int { // func wrapper() is function def, with passed in parameters, func() int is the return value
	
	w := 0
	return func() int {
		w++
		return w
	}
}


func main() {
	x := 42
	fmt.Println(x)
	// below is kindof like an anonymous function
	{
		fmt.Println(x) 
		y:= "the" // only accessible here, not down
		fmt.Println(y)
	}
	// fmt.Println(y) //outside of scope of y

	fmt.Println(increment())
	fmt.Println(increment())

	// y is at the narrow function, and anonymous function scope, rather than package level
	y := 0
	incrementy := func() int {
		y++
		return y
	}

	fmt.Println(incrementy())
	fmt.Println(incrementy())
	
	incrementw := wrapper()
	fmt.Println(incrementw())
	fmt.Println(incrementw())

	
	// anonymous function is above, has no name
}
