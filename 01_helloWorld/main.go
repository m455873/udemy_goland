package main

import "fmt"

func main() {
	//fmt.Println("hello world!")
	//fmt.Println(42)
	//fmt.Printf("%d - %b\n", 42, 42) // printf takes as many values as you want after format string
	// variadic

	fmt.Printf("%d - %b - %x\n", 42, 42, 42) // other format verbs
	fmt.Printf("%d - %b - %#x\n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X\n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#x\n", 42, 42, 42)

	// so, don't need brackets around control statement, do need curlies
	// variable creation and initialization to value use :=, where as just assignment is =
	// no while loops?  just use for in the same format as while
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q\n", i, i, i, i) //%q is for utf-8
	}

}
