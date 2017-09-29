package main

import "fmt"

const metersToYards float64 = 1.09361 // this const is typed
// interestingly, the type is defined after const [variable name] [type]

func main() {
	a := 42

	fmt.Println("a - ",a)
	fmt.Println("the address of a is:",&a) // reference to a

	var meters float64 // create variable, no value yet (default val?)
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters) // scan prompts and takes input from STDIN, passed in ref to local variable
	yards := meters * metersToYards
	fmt.Println(meters," meters are:", yards, " yards")
}
